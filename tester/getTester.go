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

const getUrl string = "http://127.0.0.1:8080/lookup/vpn/"
const getUrl2 string = "http://127.0.0.1:8080/lookup/ip/"

//const getUrl string = "http://127.0.0.1:8080/"

func requestTest(requestCount int32, requestUrl string) float64 {
	requestChannel := make(chan string)
	responseChannel := make(chan string)
	errorChannel := make(chan string)
	var execTime float64
	goRoutineCreateCount := 10

	// requestChannel 을 통해서 전달받은 ip 를 이용하여 lookup api 호출
	// 호출결과는 정상과 오류를 분리하여 결과 채널에 추가후 대기
	for i := 0; i < goRoutineCreateCount; i++ {
		go func(goRoutineNumber int, reqch <-chan string, respch chan<- string, errch chan<- string) {
			totalTimer := tool.NewTimer()
			var totalRequest int
			var totalTime float64
			for {
				requestString, more := <-reqch
				if more {
					timer := tool.NewTimer()
					resp, err := http.Get(requestUrl + requestString)
					//resp, err := http.Get(getUrl)
					totalRequest += 1
					if err != nil {
						errch <- fmt.Sprintf("[gonum:%d] %s > %v", goRoutineNumber, err.Error(), timer.Stop())
					} else {
						respBody, _ := ioutil.ReadAll(resp.Body)
						//fmt.Printf("[gonum:%d] %s > %.4f ms\n", goRoutineNumber, respBody[:len(respBody)-1], float64(timer.Stop())/float64(int64(time.Millisecond)))
						respch <- fmt.Sprintf("[gonum:%d] %s > %v", goRoutineNumber, respBody[:len(respBody)-1], timer.Stop())
						if err := resp.Body.Close(); err != nil {
							fmt.Println(err)
						}
						totalTime += timer.Stop()
					}

				} else {
					if execTime < totalTimer.Stop() {
						execTime = totalTimer.Stop()
						//fmt.Printf("check set exectime > %.4f ms\n", float64(execTime) / float64(time.Millisecond))
						//fmt.Printf("sum set exectime > %.4f ms\n", float64(totalTime) / float64(time.Millisecond))
					}
					return
				}
			}
		}(i+1, requestChannel, responseChannel, errorChannel)
	}

	go func() {
		// request channel 종료, 하위 고루틴 종료 처리
		defer close(requestChannel)
		for i := int32(0); i < requestCount; i++ {
			requestChannel <- randomdata.IpV4Address()
			//time.Sleep(1 * time.Millisecond)
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
	return execTime
}

func main() {
	//requestRange := []int32{100, 1000, 5000, 10000, 20000, 30000, 50000}
	requestRange := []int32{20}
	for _, i := range requestRange {
		execTime := requestTest(i, getUrl)
		fmt.Printf("%s request count > %d  time > %.4f / %.4f = %.4f ms\n", getUrl, i, execTime, float64(i), execTime/float64(time.Millisecond)/float64(i))
		execTime = requestTest(i, getUrl2)
		fmt.Printf("%s request count > %d  time > %.4f ms\n", getUrl2, i, execTime/float64(time.Millisecond)/float64(i))
	}
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
