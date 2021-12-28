package fly

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/sdl"
)

type FlyItem struct {
	posX float64
	posY float64
	dX   float64
	dY   float64
}

type FlyWidget struct {
	widgets.BaseWidget
	item *FlyItem
}

func NewFlyWidget() *FlyWidget {
	newWelcomeWidget := new(FlyWidget)

	newFlyItem := new(FlyItem)
	newFlyItem.dX = 2.5
	newFlyItem.dY = 7.02
	newWelcomeWidget.item = newFlyItem

	return newWelcomeWidget
}

func (self *FlyWidget) ProcessEvent(e *evt.Event) {
}

func (self *FlyWidget) SetRect(x int, y int, width int, height int) {
	self.BaseWidget.SetRect(x, y, width, height)
	/* Set start FLY position */
	newFlyItem := self.item
	newFlyItem.posX = float64(x + width/2)
	newFlyItem.posY = float64(y + height/2)
}

func (self *FlyWidget) Render() {

	self.BaseWidget.Render()

	mainRenderer := ctx.GetRenderer()

	newFlyItem := self.item
	newFlyItem.posX = newFlyItem.posX + newFlyItem.dX
	newFlyItem.posY = newFlyItem.posY + newFlyItem.dY

	maxX := self.X + self.Width
	if newFlyItem.posX < float64(self.X) || newFlyItem.posX > float64(maxX) {
		newFlyItem.dX = -newFlyItem.dX
	}
	maxY := self.Y + self.Height
	if newFlyItem.posY < float64(self.Y) || newFlyItem.posY > float64(maxY) {
		newFlyItem.dY = -newFlyItem.dY
	}

	mainRenderer.SetDrawColor(0xFF, 0, 0, 0xFF)

	rect := sdl.Rect{int32(newFlyItem.posX) - 5, int32(newFlyItem.posY) - 5, 10, 10}
	mainRenderer.FillRect(&rect)

}
