package clock

import (
	"fmt"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"time"
)

type ClockWidget struct {
	x int
	y int
}

func NewClockWidget() *ClockWidget {
	return new(ClockWidget)
}

func (self *ClockWidget) ProcessEvent(e *evt.Event) {
}

func (self *ClockWidget) Render() {

	nowTime := time.Now()

	backgroundWidget := widgets.NewRectangleWidget()
	backgroundWidget.SetRect(300, 300, 450, 650)
	backgroundWidget.SetColor(255, 255, 255, 192)
	backgroundWidget.Render()

	clockTextWidget := widgets.NewTextWidget("", 36)
	clockTextWidget.SetRect(300, 300, 100, 100)
	clockTextWidget.SetColor(255, 0, 0, 0)
	clockTextWidget.SetText(fmt.Sprintf("%02d:%02d", nowTime.Hour(), nowTime.Minute()))
	clockTextWidget.Render()
	clockTextWidget.Destroy()

}
