package smarthome

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/smarthome/bank"
	"github.com/bringdesk/bringdesk/smarthome/clock"
	"github.com/bringdesk/bringdesk/smarthome/debug"
	"github.com/bringdesk/bringdesk/smarthome/gismeteo"
	"github.com/bringdesk/bringdesk/smarthome/timer"
	"github.com/bringdesk/bringdesk/smarthome/todoist"
	"github.com/bringdesk/bringdesk/smarthome/welcome"
	"github.com/bringdesk/bringdesk/widgets"
	"path"
)

type DashboardWidget struct {
	widgets.IWidget
	widget           widgets.IWidget
	backgroundWidget widgets.IWidget
}

func NewMainWidget() *DashboardWidget {

	newMainWidget := new(DashboardWidget)

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

	/* Initialize main screen */
	mainWidgetGroup := widgets.NewWidgetGroup()

	gismeteoWidget := gismeteo.NewGismeteoWidget()
	gismeteoWidget.SetRect(100, 100, 240, 320)
	//gismeteoWidget.SetUpdateInterval(10 * time.Minute)
	mainWidgetGroup.RegisterWidget(gismeteoWidget)

	welcomeWidget := welcome.NewWelcomeWidget()
	welcomeWidget.SetRect(100+240+10, 100, 240, 320)
	mainWidgetGroup.RegisterWidget(welcomeWidget)

	newTimerWidget := timer.NewTimerWidget()
	newTimerWidget.SetRect(100+240+10+240+10, 100, 240, 320)
	mainWidgetGroup.RegisterWidget(newTimerWidget)

	newBankWidget := bank.NewBankWidget()
	newBankWidget.SetRect(100+240+10+240+10+240+10, 100, 240, 320)
	mainWidgetGroup.RegisterWidget(newBankWidget)

	clockWidget := clock.NewClockWidget()
	clockWidget.SetRect(100+240+10+240+10+240+10+240+10, 100, 240, 320)
	mainWidgetGroup.RegisterWidget(clockWidget)

	/* Debug widget */
	debugWidget := debug.NewDebugWidget()
	debugWidget.SetRect(100+240+10+240+10+240+10+240+10+240+10, 100, 240, 320)
	mainWidgetGroup.RegisterWidget(debugWidget)

	/* Todoist Widget */
	todoistWidget := todoist.NewTodoistWidget()
	todoistWidget.SetRect(100, 100+320+10, 240+10+240+10+240, 320)
	mainWidgetGroup.RegisterWidget(todoistWidget)

	/* Save */
	newMainWidget.widget = mainWidgetGroup

	return newMainWidget
}

func (self *DashboardWidget) ProcessEvent(e *evt.Event) {
	self.widget.ProcessEvent(e)
}

func (self *DashboardWidget) Render() {

	/* Render background */
	self.backgroundWidget.Render()

	/* Render widgets */
	self.widget.Render()
}
