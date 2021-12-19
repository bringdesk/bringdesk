package smarthome

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/smarthome/bank"
	"github.com/bringdesk/bringdesk/smarthome/clock"
	"github.com/bringdesk/bringdesk/smarthome/gismeteo"
	"github.com/bringdesk/bringdesk/smarthome/timer"
	"github.com/bringdesk/bringdesk/smarthome/welcome"
	"github.com/bringdesk/bringdesk/widgets"
	"path"
)

type MainWidget struct {
	widgets.IWidget
	widget widgets.IWidget
}

func NewMainWidget() *MainWidget {

	mainDir := ctx.GetBaseDir()
	mainSkin := ctx.GetSkin()

	newMainWidget := new(MainWidget)

	/* Initialize main screen */
	mainWidgetGroup := widgets.NewWidgetGroup()

	imageName := mainSkin.GetBgImage()
	newPath := path.Join(mainDir, "resources", "wallpapers", imageName)
	backgroundWidget := widgets.NewImageWidget(newPath)
	mainWidgetGroup.RegisterWidget(backgroundWidget)

	gismeteoWidget := gismeteo.NewGismeteoWidget()
	//gismeteoWidget.SetUpdateInterval(10 * time.Minute)
	mainWidgetGroup.RegisterWidget(gismeteoWidget)

	welcomeWidget := welcome.NewWelcomeWidget()
	mainWidgetGroup.RegisterWidget(welcomeWidget)

	newTimerWidget := timer.NewTimerWidget()
	mainWidgetGroup.RegisterWidget(newTimerWidget)

	newBankWidget := bank.NewBankWidget()
	mainWidgetGroup.RegisterWidget(newBankWidget)

	clockWidget := clock.NewClockWidget()
	//clockWidget.SetRect()
	mainWidgetGroup.RegisterWidget(clockWidget)

	/* Save */
	newMainWidget.widget = mainWidgetGroup

	return newMainWidget
}

func (self *MainWidget) Render() {
	self.widget.Render()
}
