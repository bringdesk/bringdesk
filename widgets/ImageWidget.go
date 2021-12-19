package widgets

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type ImageWidget struct {
	path    string
	surface *sdl.Surface
	texture *sdl.Texture
}

func NewImageWidget(path string) *ImageWidget {
	iw := new(ImageWidget)
	iw.path = path

	surface, err1 := img.Load(path)
	if err1 != nil {
		panic(err1)
	}
	iw.surface = surface

	mainRenderer := ctx.GetRenderer()

	//newTexture :=
	iw.texture, _ = mainRenderer.CreateTextureFromSurface(surface)
	//defer newTexture.Destroy()

	return iw

}

func (self *ImageWidget) ProcessEvent(e *evt.Event) {
}

func (self *ImageWidget) Render() {

	mainRect := ctx.GetRect()
	mainRenderer := ctx.GetRenderer()

	newRect := &sdl.Rect{0, 0, mainRect.W, mainRect.H}

	mainRenderer.Clear()
	mainRenderer.Copy(self.texture, nil, newRect) //SDL_RenderCopy(renderer, bitmapTex, NULL, NULL)

}

func (self *ImageWidget) Destroy() {
	/* Release texture */
	err1 := self.texture.Destroy()
	log.Printf("erorr = %#v", err1)
	/* Release surface */
	self.surface.Free()
}
