package widgets

type WidgetGroup struct {
	widgets []IWidget
}

func NewWidgetGroup() *WidgetGroup {
	wg := new(WidgetGroup)
	return wg
}

func (self *WidgetGroup) RegisterWidget(w IWidget) {
	self.widgets = append(self.widgets, w)
}

func (self *WidgetGroup) Render() {
	for _, w := range self.widgets {
		w.Render()
	}
}
