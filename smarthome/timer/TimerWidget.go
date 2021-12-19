package timer

import (
	"fmt"
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/mix"
	"path"
	"time"
)

type TimerWidget struct {
	widgets.IWidget
	startTime time.Time
	stopTime  time.Time
	mainSound *mix.Chunk
}

func NewTimerWidget() *TimerWidget {
	newTimerWidget := new(TimerWidget)
	newTimerWidget.startTime = time.Now()
	newTimerWidget.stopTime = time.Now()

	mainDir := ctx.GetBaseDir()
	newPath := path.Join(mainDir, "resources", "sounds", "Belligerent.wav")
	chunk, err1 := mix.LoadWAV(newPath)
	if err1 != nil {
		panic(err1)
	}
	newTimerWidget.mainSound = chunk

	return newTimerWidget
}

func (self *TimerWidget) ProcessEvent(e *evt.Event) {
}

func (self *TimerWidget) Render() {

	//mainRenderer := ctx.GetRenderer()

	elapse := time.Until(self.stopTime)
	//log.Printf("Timer: elapse = %q", elapse)
	var waitTime time.Duration = 1 * time.Second
	if elapse > waitTime {

		newOut := widgets.NewTextWidget("", 35)
		newOut.SetRect(500, 500, 100, 100)

		newText := fmt.Sprintf("%02d:%02d.%01d",
			int(elapse.Minutes())%60,
			int(elapse.Seconds())%60,
			(elapse.Milliseconds()%1000)/100,
		)
		newOut.SetText(newText)
		newOut.Render()
		newOut.Destroy()

	} else {
		/**/
		self.mainSound.Play(-1, 0)
		/**/
		self.startTime = time.Now()
		newStopTime := time.Now()
		newStopTime = newStopTime.Add(5 * time.Minute)
		self.stopTime = newStopTime
	}

}

func (self *TimerWidget) Destroy() {
	self.mainSound.Free()
}
