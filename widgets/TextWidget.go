package widgets

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
)

type TextWidget struct {
	IWidget          /* Widget interface is implementing */
	body     string    /* Widget content                   */
	font     *ttf.Font /* Widget font                      */
	color    sdl.Color
	height   int32
	width    int32
	x        int32
	y        int32

	fntAlias string
	fntSize  int
}

func NewTextWidget(fontAlias string, fontSize int) *TextWidget {
	newTextWidget := new(TextWidget)
	newTextWidget.fntAlias = fontAlias
	newTextWidget.fntSize = fontSize
	return newTextWidget
}

func (self *TextWidget) SetColor(r byte, g byte, b byte, a byte) {
	self.color.R = r
	self.color.G = g
	self.color.B = b
	self.color.A = a
}

func (self *TextWidget) Render() {

	mainFontManager := ctx.GetFontManager()
	mainRenderer := ctx.GetRenderer()

	//mainRenderer.SetClipRect()

	/**/
	font, _ := mainFontManager.Acquire(self.fntAlias, self.fntSize)

	/**/
	surface, err1 := font.RenderUTF8Blended(self.body, self.color)
	if err1 != nil {
		log.Printf("RenderUTF8Blended: err = %#v", err1)
	}
	texture, _ := mainRenderer.CreateTextureFromSurface(surface)

	newPosition := sdl.Rect{self.x, self.y, surface.W, surface.H}

	mainRenderer.Copy(texture, nil, &newPosition)

	/* Release resources */
	mainFontManager.Release(font)

	surface.Free()
	texture.Destroy()

}

func (self *TextWidget) SetText(body string) {
	self.body = body
}

func (self *TextWidget) SetRect(x int, y int, width int, height int) {
	self.x = int32(x)
	self.y = int32(y)
	self.width = int32(width)
	self.height = int32(height)
}

func (self *TextWidget) Destroy() {
	//self.font.Close()
}
