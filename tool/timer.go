package tool

import (
	"time"
)

type Timer struct {
	startTime_ time.Time
}

func (timer *Timer) Start() {
	timer.startTime_ = time.Now()
}

func (timer *Timer) Stop() time.Duration {
	if timer.startTime_.Unix() != int64(0) {
		return time.Since(timer.startTime_)
	} else {
		return 0
	}
}

func NewTimer() *Timer {
	timer := new(Timer)
	timer.Start()
	return timer
}
