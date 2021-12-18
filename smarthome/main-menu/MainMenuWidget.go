package main_menu

import (
	"github.com/bringdesk/bringdesk/smarthome/clock"
	"github.com/bringdesk/bringdesk/widgets"
	"log"
)

type MenuItem struct {
	summary string
	widget  widgets.IWidget
}

type MainMenu struct {
	items []*MenuItem
}

func NewMainMenu() *MainMenu {
	newMainMenu := new(MainMenu)
	newItem1 := new(MenuItem)
	newItem1.summary = "Clock"
	newItem1.widget = clock.NewClockWidget()
	newMainMenu.items = append(newMainMenu.items, newItem1)
	return newMainMenu
}

func (self *MainMenu) Render() {
	for _, item := range self.items {
		log.Printf("Menu item draw: item = %#v", item)
	}
}
