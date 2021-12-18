package smarthome

import (
	"github.com/bringdesk/bringdesk/smarthome/welcome"
	"github.com/bringdesk/bringdesk/widgets"
)

type MainWidget struct {
	widgets.IWidget
	widget widgets.IWidget
}

func NewMainWidget() *MainWidget {
	newMainWidget := new(MainWidget)

	/* Initialize main screen */
	mainWidgetGroup := widgets.NewWidgetGroup()

	backgroundWidget := widgets.NewImageWidget("C:\\Users\\vit12\\Work\\bringdesk\\wallpapers\\pexels-cottonbro-4937197.jpg")
	mainWidgetGroup.RegisterWidget(backgroundWidget)

	welcomeMessageWidget := widgets.NewTextWidget("PublicSans", 21)
	welcomeMessageWidget.SetColor(0, 255, 0, 255)
	welcomeMessageWidget.SetText("123 Hello $ Тест")
	welcomeMessageWidget.SetRect(100, 100, 200, 200)

	mainWidgetGroup.RegisterWidget(welcomeMessageWidget)

	welcomeWidget := welcome.NewWelcomeWidget()
	mainWidgetGroup.RegisterWidget(welcomeWidget)

	/* Save */
	newMainWidget.widget = mainWidgetGroup

	return newMainWidget
}

func (self *MainWidget) Render() {
	self.widget.Render()
}
