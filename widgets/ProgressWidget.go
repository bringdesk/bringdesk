package widgets

type ProgressWidget struct {
	current int64
	total   int64
}

func NewProgressWidget() *ProgressWidget {
	newProgressWidget := new(ProgressWidget)
	newProgressWidget.current = 0
	newProgressWidget.total = 100
	return newProgressWidget
}

func (self *ProgressWidget) Render() {

	//mainRenderer  := ctx.GetRenderer()

	//newColor := sdl.Color{}
	//newColor.R = 255

	//gfx.RoundedBoxColor(mainRenderer, 10, 20, 30, 40, 10, newColor)

}
