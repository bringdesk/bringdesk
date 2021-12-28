package smarthome

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/layout"
	"github.com/bringdesk/bringdesk/smarthome/bank"
	"github.com/bringdesk/bringdesk/smarthome/clock"
	"github.com/bringdesk/bringdesk/smarthome/debug"
	"github.com/bringdesk/bringdesk/smarthome/fly"
	"github.com/bringdesk/bringdesk/smarthome/news"
	"github.com/bringdesk/bringdesk/smarthome/openweathermap"
	security_system "github.com/bringdesk/bringdesk/smarthome/security-system"
	"github.com/bringdesk/bringdesk/smarthome/timer"
	"github.com/bringdesk/bringdesk/smarthome/todoist"
	"github.com/bringdesk/bringdesk/widgets"
	"log"
	"path"
)

type DeskMode int

const (
	DeskModeActive = DeskMode(1)
	DeskModeManage = DeskMode(2)
)

type DashboardWidget struct {
	widgets.IWidget
	widget                 *widgets.WidgetGroup
	backgroundWidget       widgets.IWidget
	state                  DeskMode
	dashboardWidgetFactory *DashboardWidgetFactory
}

func (self *DashboardWidget) registerWidget(w widgets.IWidget) {
	self.widget.RegisterWidget(w)
}

func NewMainWidget() *DashboardWidget {

	newMainWidget := new(DashboardWidget)
	newMainWidget.state = DeskModeActive
	newMainWidget.dashboardWidgetFactory = NewDashboardWidgetFactory()

	/* Initialize skin and backgrounds */
	mainDir := ctx.GetBaseDir()
	mainSkin := ctx.GetSkin()
	imageName := mainSkin.GetBgImage()
	if imageName != "" {
		newPath := path.Join(mainDir, "resources", "wallpapers", imageName)
		newMainWidget.backgroundWidget = widgets.NewImageWidget(newPath)
	} else {
		// TODO - make solid color rectangle widget ...
	}

	/* Initialize widget group */
	mainWidgetGroup := widgets.NewWidgetGroup()
	newMainWidget.widget = mainWidgetGroup

	/* Initialize common widgets */
	newMainWidget.initializeCommonWidgets()

	return newMainWidget
}

func (self *DashboardWidget) ProcessEvent(e *evt.Event) {
	self.widget.ProcessEvent(e)
}

func (self *DashboardWidget) Render() {

	self.backgroundWidget.Render()

	/* Render background */
	if self.state == DeskModeActive {
		/* Render widgets */
		self.widget.Render()
	} else if self.state == DeskModeManage {
		// TODO - draw only widget rects ...

	} else {
		log.Printf("Unknown state. Switch back on `DeskModeActive`")
		self.state = DeskModeActive
	}

}

func (self *DashboardWidget) initializeCommonWidgets() {

	mainRect := ctx.GetRect()

	self.initializeComponents()

	basicLayout := layout.NewBasicGridLayout()
	basicLayout.SetSize(int(mainRect.W), int(mainRect.H))

	weatherWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("OpenWeatherMapWidget")
	basicLayout.UpdatePos(0, 0, weatherWidget)
	self.registerWidget(weatherWidget)

	exampleWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("FlyWidget")
	basicLayout.UpdatePos(0, 1, exampleWidget)
	self.registerWidget(exampleWidget)

	newTimerWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("TimerWidget")
	basicLayout.UpdatePos(0, 2, newTimerWidget)
	self.registerWidget(newTimerWidget)

	newBankWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("BankWidget")
	basicLayout.UpdatePos(0, 3, newBankWidget)
	self.registerWidget(newBankWidget)

	clockWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("ClockWidget")
	basicLayout.UpdatePos(0, 4, clockWidget)
	self.registerWidget(clockWidget)

	/* Debug widget */
	debugWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("DebugWidget")
	basicLayout.UpdatePos(0, 5, debugWidget)
	self.registerWidget(debugWidget)

	/* Todoist Widget */
	todoistWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("TodoistWidget")
	basicLayout.UpdatePos(1, 0, todoistWidget)
	self.registerWidget(todoistWidget)

	/* Secure camera */
	secureCameraWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("SecuritySystemWidget")
	basicLayout.UpdatePos(1, 1, secureCameraWidget)
	self.registerWidget(secureCameraWidget)

	/* News widget */
	newsWidget, _ := self.dashboardWidgetFactory.CreateWidgetByName("NewsWidget")
	basicLayout.UpdatePos(2, 0, newsWidget)
	self.registerWidget(newsWidget)

}

func (self *DashboardWidget) initializeComponents() {
	/* Step 1. */
	self.dashboardWidgetFactory.RegisterWidget(
		"OpenWeatherMapWidget",
		func() widgets.IWidget {
			return openweathermap.NewOpenWeatherMapWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"FlyWidget",
		func() widgets.IWidget {
			return fly.NewFlyWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"TimerWidget",
		func() widgets.IWidget {
			return timer.NewTimerWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"BankWidget",
		func() widgets.IWidget {
			return bank.NewBankWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"ClockWidget",
		func() widgets.IWidget {
			return clock.NewClockWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"DebugWidget",
		func() widgets.IWidget {
			return debug.NewDebugWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"TodoistWidget",
		func() widgets.IWidget {
			return todoist.NewTodoistWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"SecuritySystemWidget",
		func() widgets.IWidget {
			return security_system.NewSecuritySystemWidget()
		},
	)
	self.dashboardWidgetFactory.RegisterWidget(
		"NewsWidget",
		func() widgets.IWidget {
			return news.NewNewsWidget()
		},
	)
}
