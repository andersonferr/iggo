package ixlib

import (
	"unsafe"

	"github.com/andersonferr/iggo/backend"
)

var xWindow2Handler map[xWindow]*Handler

func init() {
	xWindow2Handler = make(map[xWindow]*Handler)
}

var eventBuffer struct {
	buffer [4 * 1024]xEvent
	size   int
	curr   int
}

//NextEvent return the next event or nil
func (env *Environment) NextEvent(ev *backend.Event) {
	if eventBuffer.curr == eventBuffer.size {
		eventBuffer.curr = 0
		eventBuffer.size = readEvents(env.display, eventBuffer.buffer[:])

		if eventBuffer.size == 0 {
			return
		}
	}

	xEventType := *(*xInt)(unsafe.Pointer(&eventBuffer.buffer[eventBuffer.curr]))
	switch xEventType {
	case xExpose:
		xev := (*xExposeEvent)(unsafe.Pointer(&eventBuffer.buffer[eventBuffer.curr]))
		w := xev.window
		h := xWindow2Handler[w]
		ev.Type = backend.EventTypeExpose
		ev.Handler = h

	case xClientMessage:
		xev := (*xClientMessageEvent)(unsafe.Pointer(&eventBuffer.buffer[eventBuffer.curr]))
		w := xev.window
		h := xWindow2Handler[w]
		ev.Type = backend.EventTypeClose
		ev.Handler = h

	case xConfigureNotify:
		xev := (*xConfigureEvent)(unsafe.Pointer(&eventBuffer.buffer[eventBuffer.curr]))
		w := xev.window
		h := xWindow2Handler[w]

		ev.Type = backend.EventTypeResize
		ev.Handler = h
		ev.Width = int(xev.width)
		ev.Height = int(xev.height)
	}

	eventBuffer.curr++
}
