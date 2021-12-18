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

// StateSave save timer on disk
func StateSave() {

}

// StateRestore restore timer state from disk
func StateRestore() {

}
