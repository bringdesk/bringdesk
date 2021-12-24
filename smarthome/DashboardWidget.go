package smarthome

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/smarthome/bank"
	"github.com/bringdesk/bringdesk/smarthome/clock"
	"github.com/bringdesk/bringdesk/smarthome/debug"
	"github.com/bringdesk/bringdesk/smarthome/news"
	"github.com/bringdesk/bringdesk/smarthome/openweathermap"
	security_system "github.com/bringdesk/bringdesk/smarthome/security-system"
	"github.com/bringdesk/bringdesk/smarthome/timer"
	"github.com/bringdesk/bringdesk/smarthome/todoist"
	"github.com/bringdesk/bringdesk/smarthome/welcome"
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
	widget           *widgets.WidgetGroup
	backgroundWidget widgets.IWidget
	state            DeskMode
}

func (self *DashboardWidget) registerWidget(w widgets.IWidget) {
	self.widget.RegisterWidget(w)
}

func NewMainWidget() *DashboardWidget {

	newMainWidget := new(DashboardWidget)
	newMainWidget.state = DeskModeActive

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

	openweathermapWidget := openweathermap.NewOpenWeatherMapWidget()
	openweathermapWidget.SetRect(100, 100, 240, 320)
	self.registerWidget(openweathermapWidget)

	welcomeWidget := welcome.NewWelcomeWidget()
	welcomeWidget.SetRect(100+240+10, 100, 240, 320)
	self.registerWidget(welcomeWidget)

	newTimerWidget := timer.NewTimerWidget()
	newTimerWidget.SetRect(100+240+10+240+10, 100, 240, 320)
	self.registerWidget(newTimerWidget)

	newBankWidget := bank.NewBankWidget()
	newBankWidget.SetRect(100+240+10+240+10+240+10, 100, 240, 320)
	self.registerWidget(newBankWidget)

	clockWidget := clock.NewClockWidget()
	clockWidget.SetRect(100+240+10+240+10+240+10+240+10, 100, 240, 320)
	self.registerWidget(clockWidget)

	/* Debug widget */
	debugWidget := debug.NewDebugWidget()
	debugWidget.SetRect(100+240+10+240+10+240+10+240+10+240+10, 100, 240, 320)
	self.registerWidget(debugWidget)

	/* Todoist Widget */
	todoistWidget := todoist.NewTodoistWidget()
	todoistWidget.SetRect(100, 100+320+10, 240+10+240+10+240, 320)
	self.registerWidget(todoistWidget)

	/* Secure camera */
	secureCameraWidget := security_system.NewSecuritySystemWidget()
	secureCameraWidget.SetRect(100+240+10+240+10+240+10, 100+320+10, 240+10+240+10+240, 320)
	self.registerWidget(secureCameraWidget)

	/* News widget */
	newsWidget := news.NewNewsWidget()
	newsWidget.SetRect(100, 100+320+10+320+10, 240+10+240+10+240+10+240+10+240+10+240, 320)
	self.registerWidget(newsWidget)

}
