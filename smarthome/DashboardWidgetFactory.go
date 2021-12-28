package smarthome

import "github.com/bringdesk/bringdesk/widgets"

type DashboardWidgetFactory struct {
	widgetMap map[string]func() widgets.IWidget
}

func NewDashboardWidgetFactory() *DashboardWidgetFactory {
	newDashboardWidgetFactory := &DashboardWidgetFactory{
		widgetMap: make(map[string]func() widgets.IWidget),
	}
	return newDashboardWidgetFactory
}

func (self *DashboardWidgetFactory) RegisterWidget(widgetName string, create func() widgets.IWidget) {
	self.widgetMap[widgetName] = create
}

func (self *DashboardWidgetFactory) CreateWidgetByName(widgetName string) (widgets.IWidget, error) {
	widgetCreate := self.widgetMap[widgetName]
	return widgetCreate(), nil
}
