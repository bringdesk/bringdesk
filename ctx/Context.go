package ctx

import (
	"github.com/bringdesk/bringdesk/skin"
	"github.com/veandco/go-sdl2/sdl"
)

var mainWindow *sdl.Window
var mainSurface *sdl.Surface
var mainRenderer *sdl.Renderer
var mainRect *sdl.Rect
var mainDir string
var mainSkin *skin.Skin

func GetWindow() *sdl.Window {
	return mainWindow
}

func SetWindow(window *sdl.Window) {
	mainWindow = window
}

func SetSurface(surface *sdl.Surface) {
	mainSurface = surface
}

func GetSurface() *sdl.Surface {
	return mainSurface
}

func SetRenderer(renderer *sdl.Renderer) {
	mainRenderer = renderer
}

func GetRenderer() *sdl.Renderer {
	return mainRenderer
}

func GetRect() *sdl.Rect {
	return mainRect
}

func SetRect(rect *sdl.Rect) {
	mainRect = rect
}

func GetBaseDir() string {
	return mainDir
}

func SetBaseDir(baseDir string) {
	mainDir = baseDir
}

func SetSkin(newSkin *skin.Skin) {
	mainSkin = newSkin
}

func GetSkin() *skin.Skin {
	return mainSkin
}
