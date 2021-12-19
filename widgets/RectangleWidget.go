package widgets

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type RectangleWidget struct {
	r      uint8
	a      uint8
	b      uint8
	g      uint8
	x      int
	y      int
	height int
	width  int
}

func NewRectangleWidget() *RectangleWidget {
	return new(RectangleWidget)
}

func (self *RectangleWidget) SetColor(r, g, b, a uint8) {
	self.r = r
	self.g = g
	self.b = b
	self.a = a
}

func (self *RectangleWidget) SetRect(x, y, width, height int) {
	self.x = x
	self.y = y
	self.width = width
	self.height = height
}

func (self *RectangleWidget) Render() {

	mainRenderer := ctx.GetRenderer()

	/* Clear screen */
	newRect := sdl.Rect{int32(self.x), int32(self.y), int32(self.width), int32(self.height)}
	mainRenderer.SetDrawColor(self.r, self.g, self.b, self.a)
	mainRenderer.FillRect(&newRect)

}
