package timer

import "time"

type TimerWidget struct {
	startTime time.Time
	stopTime  time.Time
}

func NewTimerWidget() *TimerWidget {
	newTimerWidget := new(TimerWidget)
	return newTimerWidget
}
