package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sampler/tool"
	"time"

	"github.com/Pallinder/go-randomdata"
)

//	100개의 고루틴 생성후 10000개의 리퀘스트를 채널을 통해 고루틴에 요청.

//const getUrl string = "http://127.0.0.1:8080/lookup/ip/"
const getUrl string = "http://127.0.0.1:8080/"

func main() {

	requestChannel := make(chan string)
	responseChannel := make(chan string)
	errorChannel := make(chan string)

	// requestChannel 을 통해서 전달받은 ip 를 이용하여 lookup api 호출
	// 호출결과는 정상과 오류를 분리하여 결과 채널에 추가후 대기
	for i := 0; i < 1; i++ {
		go func(goRoutineNumber int, reqch <-chan string, respch chan<- string, errch chan<- string) {
			totalTimer := tool.NewTimer()
			var totalRequest int
			for {
				requestString, more := <-reqch
				if more {
					timer := tool.NewTimer()
					//resp, err := http.Get(getUrl + requestString)
					resp, err := http.Get(getUrl)
					totalRequest += 1
					if err != nil {
						errch <- fmt.Sprintf("[gonum:%d] %s > %v", goRoutineNumber, err.Error(), timer.Stop())
					} else {
						respBody, _ := ioutil.ReadAll(resp.Body)
						fmt.Printf("[gonum:%d] %s > %v\n", goRoutineNumber, respBody[:len(respBody)-1], timer.Stop())
						respch <- fmt.Sprintf("[gonum:%d] %s > %v", goRoutineNumber, respBody[:len(respBody)-1], timer.Stop())
					}
					resp.Body.Close()
				} else {
					fmt.Println(requestString)
					fmt.Printf("goroutine[%d] finish > total request %d ea [%v]\n", goRoutineNumber, totalRequest, totalTimer.Stop())
					return
				}
			}
		}(i+1, requestChannel, responseChannel, errorChannel)
	}

	go func() {
		// request channel 종료, 하위 고루틴 종료 처리
		defer close(requestChannel)
		for i := 0; i < 1000; i++ {
			requestChannel <- randomdata.IpV4Address()
			time.Sleep(1 * time.Millisecond)
		}
	}()
	var responses []string
loop:
	for {
		select {
		case r := <-responseChannel:
			responses = append(responses, r)
		case e := <-errorChannel:
			responses = append(responses, e)
			//fmt.Errorf("%s\n", e)
		case <-time.After(500 * time.Millisecond):
			break loop
		}
	}
	// for _, val := range responses {
	// 	fmt.Println(val)
	// }
}

//	lookup api 를 호출하는 테스트
//	최초로 실행된 고루틴들을 통해 get request 를 날리고 결과를 channel 에추가
//	main 에서 채널을 읽어 출력.
/**
func main() {
	url := "http://127.0.0.1:8080/lookup/ip/"

	ch := make(chan string)
	strch := make(chan string)
	for i := 0; i < 1000; i++ {
		go func(get_url string) {
			timer := tool.NewTimer()
			resp, err := http.Get(get_url + randomdata.IpV4Address())
			if err != nil {
				strch <- err.Error()
			} else {
				respBody, _ := ioutil.ReadAll(resp.Body)
				resultString := fmt.Sprintf("%s > %v", respBody, timer.Stop())
				ch <- resultString
			}
		}(url)
	}
	var responses []string
loop:
	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
		case e := <-strch:
			fmt.Println("Error : ", e)
		case <-time.After(1 * time.Second):
			break loop
		}
	}
	for index, response := range responses {
		fmt.Printf("[index: %d] %s\n", index, response)
	}
}
*/
