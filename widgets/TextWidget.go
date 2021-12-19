package widgets

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"path"
)

type TextWidget struct {
	IWidget           /* Widget interface is implementing */
	body    string    /* Widget content                   */
	font    *ttf.Font /* Widget font                      */
	color   sdl.Color
	height  int32
	width   int32
	x       int32
	y       int32
}

func NewTextWidget(fontAlias string, fontSize int) *TextWidget {
	newTextWidget := new(TextWidget)

	fontAliases := make(map[string]string)
	fontAliases["PublicSans"] = "PublicSans-Regular.otf"

	fontFile := fontAliases[fontAlias]
	if fontFile == "" {
		fontFile = "PublicSans-Regular.otf"
	}

	baseDir := ctx.GetBaseDir()
	newPath := path.Join(baseDir, "resources", "fonts", fontFile)
	newFont, err2 := ttf.OpenFont(newPath, fontSize)
	if err2 != nil {
		panic(err2)
	}

	newTextWidget.font = newFont

	return newTextWidget
}

func (self *TextWidget) SetColor(r byte, g byte, b byte, a byte) {
	self.color.R = r
	self.color.G = g
	self.color.B = b
	self.color.A = a
}

func (self *TextWidget) Render() {

	mainRenderer := ctx.GetRenderer()

	surface, err1 := self.font.RenderUTF8Blended(self.body, self.color)
	if err1 != nil {
		log.Printf("RenderUTF8Blended: err = %#v", err1)
	}
	texture, _ := mainRenderer.CreateTextureFromSurface(surface)

	newPosition := sdl.Rect{self.x, self.y, surface.W, surface.H}

	mainRenderer.Copy(texture, nil, &newPosition)

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
	self.font.Close()
}
