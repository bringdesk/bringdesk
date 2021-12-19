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
	widgets.IWidget
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

	rectWidget := widgets.NewRectangleWidget()
	rectWidget.SetRect(600, 600, 100, 100)
	rectWidget.SetColor(100, 220,0, 128)
	rectWidget.Render()

	msgWidget := widgets.NewTextWidget("", 21)
	msg := fmt.Sprintf("Frame Rate %d Cache Font %d", self.FrameRate, self.FontCache)
	msgWidget.SetText(msg)
	msgWidget.SetColor(100, 0,0, 128)
	msgWidget.SetRect(600, 600, 100, 100)
	msgWidget.Render()
	msgWidget.Destroy()

}

func (self *DebugWidget) ProcessEvent(evt *evt.Event) {

}