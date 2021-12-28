package widgets

import "github.com/bringdesk/bringdesk/evt"

type IWidget interface {
	Render()
	ProcessEvent(e *evt.Event)
	SetRect(x int, y int, width int, height int)
}
