package main

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Application struct {
	running    bool
	mainWidget widgets.IWidget
}

func NewApplication() *Application {
	return &Application{}
}

func (self *Application) Run() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	displayCount, _ := sdl.GetNumVideoDisplays()
	log.Printf("System contain %d display(s)", displayCount)

	var rects []sdl.Rect
	for displayIndex := 0; displayIndex < displayCount; displayIndex++ {
		rect, _ := sdl.GetDisplayBounds(displayIndex)
		rects = append(rects, rect)
		log.Printf("Display %d bounds is %#v", displayIndex, rect)
	}

	var mainRect sdl.Rect = rects[1]
	window, err := sdl.CreateWindow("test", mainRect.X, mainRect.Y, mainRect.W, mainRect.H, sdl.WINDOW_SHOWN|
		sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	ctx.SetWindow(window)

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	ctx.SetSurface(surface)

	/* Initialize main screen */
	self.mainWidget = widgets.NewWelcomeWidget()

	var rate float64 = 1000 * 1.0 / 25
	log.Printf("Frame rate %.03f", rate)

	self.running = true
	for self.running {

		/* Start events processing */
		var event sdl.Event
		for {
			event = sdl.WaitEventTimeout(int(rate))
			if event == nil {
				break
			}
			e := evt.NewEventFromSDL(event)
			if e.EvType != evt.EventTypeNone {
				self.ProcessEvent(e)
			}
		}

		/* Render main scene */
		self.mainWidget.Render()

	}
}

func (self *Application) ProcessEvent(e *evt.Event) {
	if e.EvType == evt.EventTypeKeyboard {

	} else if e.EvType == evt.EventTypeMouse {
		log.Printf(
			"Mouse click on X = %d Y = %d ",
			e.Mouse.X,
			e.Mouse.Y,
		)
	} else if e.EvType == evt.EventTypeTouch {

	} else {
		log.Printf("No idea how to handle %#v", e)
	}
}

func main() {

	app := NewApplication()
	app.Run()

}
