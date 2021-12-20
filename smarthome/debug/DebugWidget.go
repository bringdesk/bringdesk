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
	renderFrameCount int
}

func NewDebugWidget() *DebugWidget {

	newDebugWidget := new(DebugWidget)
	/* Frame rate monitor */
	go func() {
		for {
			log.Printf("Frame rate %d", newDebugWidget.renderFrameCount)
			newDebugWidget.FrameRate = newDebugWidget.renderFrameCount
			newDebugWidget.renderFrameCount = 0
			time.Sleep(1 * time.Second)
		}
	}()
	return newDebugWidget
}

func (self *DebugWidget) Render() {

	self.renderFrameCount += 1

	self.BaseWidget.Render()

	mainFontManager := ctx.GetFontManager()

	/* Show Frame Rate */
	self.renderText(0, fmt.Sprintf("Frame Rate = %d", self.FrameRate))
	self.renderText(1, fmt.Sprintf("Font use counter %d", mainFontManager.GetUseFontCount()))
	self.renderText(2, fmt.Sprintf("Font count %d", mainFontManager.GetFontCount()))
	self.renderText(3, fmt.Sprintf("Cache Font %d", mainFontManager.GetFontCacheCount()))

}

func (self *DebugWidget) ProcessEvent(evt *evt.Event) {

}

func (self *DebugWidget) renderText(i int, msg string) {
	msgWidget := widgets.NewTextWidget("", 21)
	msgWidget.SetText(msg)
	msgWidget.SetColor(100, 0, 0, 128)
	msgWidget.SetRect(self.X, self.Y+i*20, self.Width, self.Height)
	msgWidget.Render()
	msgWidget.Destroy()
}
