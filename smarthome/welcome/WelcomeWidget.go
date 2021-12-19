package welcome

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type WelcomeWidget struct {
	posX        float64
	posY        float64
	dX          float64
	dY          float64
	startButton *widgets.ButtonWidget /* Start button */
}

func NewWelcomeWidget() *WelcomeWidget {
	newWelcomeWidget := new(WelcomeWidget)

	newWelcomeWidget.startButton = widgets.NewButtonWidget()
	newWelcomeWidget.startButton.SetCallback(func() {
		log.Printf("Main button pressed ")
	})

	newWelcomeWidget.dX = 2.5
	newWelcomeWidget.dY = 7.02

	return newWelcomeWidget
}

func (self *WelcomeWidget) ProcessEvent(e *evt.Event) {
}

func (self *WelcomeWidget) Render() {

	mainRenderer := ctx.GetRenderer()

	self.posX = self.posX + self.dX
	self.posY = self.posY + self.dY

	if self.posX < 0 || self.posX > 1000 {
		self.dX = -self.dX
	}
	if self.posY < 0 || self.posY > 1000 {
		self.dY = -self.dY
	}

	mainRenderer.SetDrawColor(0xFF, 0, 0, 0xFF)

	rect := sdl.Rect{int32(self.posX), int32(self.posY), 20, 20}
	mainRenderer.FillRect(&rect)

}
