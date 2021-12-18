package ctx

import "github.com/veandco/go-sdl2/sdl"

var mainWindow *sdl.Window
var mainSurface *sdl.Surface

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
