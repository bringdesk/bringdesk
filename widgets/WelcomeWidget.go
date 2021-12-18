package widgets

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type WelcomeWidget struct {
	posX float64
	posY float64
	dX   float64
	dY   float64
}

func NewWelcomeWidget() *WelcomeWidget {
	newWelcomeWidget := new(WelcomeWidget)
	newWelcomeWidget.dX = 2.5
	newWelcomeWidget.dY = 7.02
	return newWelcomeWidget
}

func (self *WelcomeWidget) Render() {

	mainWindow := ctx.GetWindow()
	mainSurface := ctx.GetSurface()

	/* Clear screen */
	mainRect := sdl.Rect{0, 0, 2000, 2000}
	mainSurface.FillRect(&mainRect, 0x00000000)

	self.posX = self.posX + self.dX
	self.posY = self.posY + self.dY

	if self.posX < 0 || self.posX > 1000 {
		self.dX = -self.dX
	}
	if self.posY < 0 || self.posY > 1000 {
		self.dY = -self.dY
	}

	rect := sdl.Rect{int32(self.posX), int32(self.posY), 20, 20}
	mainSurface.FillRect(&rect, 0xffff0000)
	mainWindow.UpdateSurface()

}
