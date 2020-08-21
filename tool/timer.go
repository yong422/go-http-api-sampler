package tool

import (
	"time"
)

type Timer struct {
	//startTime_ time.Time
	startTime_ int64
}

func (timer *Timer) Start() {
	timer.startTime_ = time.Now().UnixNano()
}

func (timer *Timer) Stop() float64 {
	return float64(time.Now().UnixNano() - timer.startTime_)
}

func (timer *Timer) ElapsedMilliseconds() float64 {
	return timer.Stop() / float64(time.Millisecond)
}

func (timer *Timer) ElapsedSeconds() float64 {
	return timer.Stop() / float64(time.Second)
}

func NewTimer() *Timer {
	timer := new(Timer)
	timer.Start()
	return timer
}
