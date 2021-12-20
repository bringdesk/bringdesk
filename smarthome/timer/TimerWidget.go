package timer

import (
	"fmt"
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"path"
	"time"
)

type TimerWidget struct {
	widgets.BaseWidget
	startTime      time.Time
	stopTime       time.Time
	mainSound     *mix.Chunk
	activeTimer    bool
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

	if e.EvType == evt.EventTypeMouse {

		point := sdl.Point{
			X: int32(e.Mouse.X),
			Y: int32(e.Mouse.Y),
		}

		rect := sdl.Rect{
			X: int32(self.X),
			Y: int32(self.Y),
			W: int32(self.Width),
			H: int32(self.Height),
		}

		if point.InRect(&rect) {
			log.Printf("Activate Timer")
			/* Activate */
			self.activeTimer = true
			/* Setup stop point */
			self.startTime = time.Now()
			newStopTime := time.Now()
			newStopTime = newStopTime.Add(15 * time.Minute)
			self.stopTime = newStopTime
		}
	}

}

func (self *TimerWidget) Render() {

	self.BaseWidget.Render()

	newOut := widgets.NewTextWidget("", 35)
	newOut.SetRect(self.X, self.Y, self.Width, self.Height)
	var newText string = "--:--:--"
	if self.activeTimer {
		elapse := time.Until(self.stopTime)
		newText = fmt.Sprintf("%02d:%02d:%02d",
			int(elapse.Hours()),
			int(elapse.Minutes())%60,
			int(elapse.Seconds())%60,
		)
	}
	newOut.SetText(newText)
	newOut.Render()
	newOut.Destroy()

	/* Play sound and stop timer */
	if self.activeTimer {
		var currentTime time.Time = time.Now()
		if currentTime.After(self.stopTime) {
			log.Printf("Timer complete.")
			/* Play sounds */
			self.mainSound.Play(-1, 0)
			/* Stop activation */
			self.activeTimer = false
		}
	}

}

func (self *TimerWidget) Destroy() {
	self.mainSound.Free()
}
