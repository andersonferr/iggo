package backend

type EventType uint32

const (
	EventTypeNoEvent EventType = iota
	EventTypeClose
	EventTypeExpose
	EventTypeResize
)

type (
	Event struct {
		Handler Handler
		Type    EventType

		Height int
		Width  int
	}
)
