package clock

import (
	"fmt"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"time"
)

type ClockWidget struct {
	widgets.BaseWidget
}

func NewClockWidget() *ClockWidget {
	return new(ClockWidget)
}

func (self *ClockWidget) ProcessEvent(e *evt.Event) {
}

func (self *ClockWidget) Render() {

	self.BaseWidget.Render()

	nowTime := time.Now()

	clockTextWidget := widgets.NewTextWidget("", 36)
	clockTextWidget.SetRect(self.X, self.Y, self.Width, self.Height)
	clockTextWidget.SetColor(255, 0, 0, 0)
	clockTextWidget.SetText(fmt.Sprintf("%02d:%02d", nowTime.Hour(), nowTime.Minute()))
	clockTextWidget.Render()
	clockTextWidget.Destroy()

}
