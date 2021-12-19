package evt

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

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

	log.Printf("NewEventFromSDL: e = %#v", e)

	switch e.(type) {
	case *sdl.KeyboardEvent:
		newEvent := &Event{
			EvType: EventTypeKeyboard,
		}
		return newEvent

	case *sdl.MouseButtonEvent:
		mouseButtonEvent, _ := e.(*sdl.MouseButtonEvent)
		newEvent := &Event{
			EvType: EventTypeMouse,
			Mouse: EventMouse{
				X: int(mouseButtonEvent.X),
				Y: int(mouseButtonEvent.Y),
			},
		}
		return newEvent

	case *sdl.QuitEvent:
		newEvent := &Event{
			EvType: EventTypeQuit,
		}
		return newEvent
	}

	return nil
}
