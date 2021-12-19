package main

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/skin"
	"github.com/bringdesk/bringdesk/smarthome"
	"github.com/bringdesk/bringdesk/util"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"os"
	"path"
	"time"
)

type Application struct {
	running          bool
	mainWidget       widgets.IWidget
}

func NewApplication() *Application {
	return &Application{}
}

func (self *Application) Run() {

	/* Initialize SDL2 */
	err1 := sdl.Init(sdl.INIT_EVERYTHING)
	if err1 != nil {
		panic(err1)
	}
	defer sdl.Quit()

	/* Initialize SDL_img */
	err2 := img.Init(img.INIT_JPG | img.INIT_PNG)
	if err2 != nil {
		panic(err2)
	}
	defer img.Quit()

	/* Initialize SDL_ttf */
	err3 := ttf.Init()
	if err3 != nil {
		panic(err3)
	}
	defer ttf.Quit()

	/* Detect resources directory */
	workDirectory, err4 := os.Getwd()
	if err4 != nil {
		panic(err4)
	}
	ctx.SetBaseDir(workDirectory)

	/* Parse arguments */
	// TODO - os.ParseArgs...

	/* Reading configuration */
	// TODO - read configuration file ...

	/* Reading skin */
	mainSkin := skin.NewSkin()
	mainSkin.DisplayIndex = 1
	//mainSkin.SetAcentColor(100, 120, 30, 255)
	mainSkin.SetBgImage("pexels-cottonbro-4937197.jpg")
	ctx.SetSkin(mainSkin)

	/* Open audio mixer */
	err5 := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 2048)
	if err5 != nil {
		log.Printf("err1 = %#v", err5)
	}

	/* Select monitor base on skin */
	displayCount, _ := sdl.GetNumVideoDisplays()
	log.Printf("System contain %d display(s)", displayCount)

	var rects []sdl.Rect
	var mainRect sdl.Rect
	for displayIndex := 0; displayIndex < displayCount; displayIndex++ {
		rect, _ := sdl.GetDisplayBounds(displayIndex)
		rects = append(rects, rect)
		mainRect = rect
	}
	ctx.SetRect(&mainRect)

	/* Create main window */
	window, err := sdl.CreateWindow("BringDesk",
		mainRect.X, mainRect.Y,
		mainRect.W, mainRect.H,
		sdl.WINDOW_SHOWN | sdl.WINDOW_FULLSCREEN_DESKTOP,
	)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	ctx.SetWindow(window)

	/* Create main renderer */
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	ctx.SetRenderer(renderer)

	/* Create main surface */
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	ctx.SetSurface(surface)

	/* Start font manager */
	newFontManager := util.NewFontManager()
	baseDir := ctx.GetBaseDir()
	searchPath := path.Join(baseDir, "resources", "fonts")
	newFontManager.SetSearchPath(searchPath)
	ctx.SetFontManager(newFontManager)

	/* Create main view */
	self.mainWidget = smarthome.NewMainWidget()

	/* Setup frame rate */
	var rate float64 = 1000 * 1.0 / 26
	log.Printf("Frame rate %.03f", rate)

	/* Set blending mode */
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	/* Main process */
	self.running = true
	for self.running {

		startTime := time.Now()

		/* Start events processing */
		var event sdl.Event
		for {
			event = sdl.PollEvent()
			if event == nil {
				break
			}
			e := evt.NewEventFromSDL(event)
			if e != nil {
				self.ProcessEvent(e)
			}
		}

		/* Render main scene */
		self.mainWidget.Render()
		renderer.Present() //SDL_RenderPresent(renderer)
		//window.UpdateSurface()

		/* Wait */
		renderDuration := time.Since(startTime)
		var newWait uint32 = 0
		var renderDurationMs int64 = renderDuration.Milliseconds()
		if float64(renderDurationMs) < rate {
			newWait = uint32(int64(rate) - renderDurationMs)
		}
		sdl.Delay(newWait)

	}
}

func (self *Application) ProcessEvent(e *evt.Event) {

	/* Main application event processing */
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

	/* Widget application event processing */
	self.mainWidget.ProcessEvent(e)

}

func main() {
	app := NewApplication()
	app.Run()
}
