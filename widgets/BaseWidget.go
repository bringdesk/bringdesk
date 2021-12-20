package widgets

type BaseWidget struct {
	X       int    /* Widget position X */
	Y       int    /* Widget position Y */
	Width   int
	Height  int
}

func (self *BaseWidget) SetRect(x int, y int, width int, height int) {
	self.X = x
	self.Y = y
	self.Width = width
	self.Height = height
}

func (self *BaseWidget) Render() {

	rectWidget := NewRectangleWidget()
	rectWidget.SetRect(self.X, self.Y, self.Width, self.Height)
	rectWidget.SetColor(100, 220,0, 128)
	rectWidget.Render()

}