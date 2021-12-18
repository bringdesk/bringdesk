package widgets

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type RectangleWidget struct {
}

func (self *RectangleWidget) Render() {

	mainSurface := ctx.GetSurface()
	mainRect := ctx.GetRect()

	/* Clear screen */
	newRect := sdl.Rect{0, 0, mainRect.W, mainRect.H}
	mainSurface.FillRect(&newRect, 0x00000000)

}
