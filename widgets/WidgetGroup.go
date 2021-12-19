package widgets

import "github.com/bringdesk/bringdesk/evt"

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

func (self *WidgetGroup) ProcessEvent(e *evt.Event) {
	for _, w := range self.widgets {
		w.ProcessEvent(e)
	}
}

func (self *WidgetGroup) Render() {
	for _, w := range self.widgets {
		w.Render()
	}
}
