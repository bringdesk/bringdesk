package evt

import "github.com/veandco/go-sdl2/sdl"

type EventType int

const (
	EventTypeQuit     = EventType(-1)
	EventTypeNone     = EventType(0)
	EventTypeKeyboard = EventType(1)
	EventTypeMouse    = EventType(2)
	EventTypeTouch    = EventType(3)
)

type EventMouse struct {
	X int
	Y int
}

type Event struct {
	EvType EventType
	Mouse  EventMouse
}

func NewEventFromSDL(e sdl.Event) *Event {

	newEvent := &Event{}

	switch e.(type) {
	case *sdl.KeyboardEvent:
		newEvent.EvType = EventTypeKeyboard
		break
	case *sdl.MouseButtonEvent:
		mouseButtonEvent, _ := e.(*sdl.MouseButtonEvent)
		newEvent.EvType = EventTypeMouse
		newEvent.Mouse.X = int(mouseButtonEvent.X)
		newEvent.Mouse.Y = int(mouseButtonEvent.Y)
		break
		//	case *sdl.MouseMotionEvent:
		//		newEvent.EvType = EventTypeMouse
	case *sdl.QuitEvent:
		newEvent.EvType = EventTypeQuit
	}

	return newEvent
}
