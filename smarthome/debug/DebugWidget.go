package debug

import (
	"fmt"
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"log"
	"time"
)

type DebugWidget struct {
	widgets.BaseWidget
	FrameRate        int
	FontCache        int
	renderFrameCount int
}

func NewDebugWidget() *DebugWidget {
	mainFontManager := ctx.GetFontManager()
	newDebugWidget := new(DebugWidget)
	/* Frame rate monitor */
	go func() {
		for {
			log.Printf("Frame rate %d", newDebugWidget.renderFrameCount)
			newDebugWidget.FrameRate = newDebugWidget.renderFrameCount
			newDebugWidget.FontCache = mainFontManager.GetUseFontCount()
			newDebugWidget.renderFrameCount = 0
			time.Sleep(1 * time.Second)
		}
	}()
	return newDebugWidget
}

func (self *DebugWidget) Render() {

	self.renderFrameCount += 1

	self.BaseWidget.Render()

	/* Show Frame Rate */
	{
		msgWidget := widgets.NewTextWidget("", 21)
		msg := fmt.Sprintf("Frame Rate = %d",
			self.FrameRate,
		)
		msgWidget.SetText(msg)
		msgWidget.SetColor(100, 0, 0, 128)
		msgWidget.SetRect(self.X, self.Y + 0*20, self.Width, self.Height)
		msgWidget.Render()
		msgWidget.Destroy()
	}

	/* Show Cache Font Count */
	{
		msgWidget := widgets.NewTextWidget("", 21)
		msg := fmt.Sprintf("Cache Font %d",
			self.FontCache,
		)
		msgWidget.SetText(msg)
		msgWidget.SetColor(100, 0, 0, 128)
		msgWidget.SetRect(self.X, self.Y + 1*20, self.Width, self.Height)
		msgWidget.Render()
		msgWidget.Destroy()
	}

}

func (self *DebugWidget) ProcessEvent(evt *evt.Event) {

}