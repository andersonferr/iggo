package ixgb

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/andersonferr/iggo/backend"
)

// Environment implements backend.Environment for XGB.
type Environment struct {
	conn   *xgb.Conn
	screen *xproto.ScreenInfo

	mu       sync.Mutex
	handlers map[xproto.Window]*Handler
}

func init() {
	// help typechecker
	var env *Environment
	_ = backend.Environment(env)
}

// EnvName the name of envimenment.
const EnvName = "IXGB"

func init() {
	backend.DefaultEnvironmentManager.RegisterProvider(EnvName, New)
}

//New creates a new backend.Environment.
func New() backend.Environment {
	return &Environment{
		handlers: make(map[xproto.Window]*Handler),
	}
}

// CreateHandler create a new backend.Handler for this environment.
func (env *Environment) CreateHandler(width, height int) backend.Handler {
	handler := newHandler(env, width, height)
	env.put(handler.windowID, handler)
	return handler
}

// Start prepare environmento to run the GUI application.
func (env *Environment) Start() {
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)

	env.conn = X
	env.screen = screen
}

// Finish releases resources.
func (env *Environment) Finish() {
	env.conn.Close()
}

// NextEvent gets the next event.
func (env *Environment) NextEvent(event *backend.Event) {
	for {
		ev, xerr := env.conn.WaitForEvent()
		if xerr != nil {
			panic(xerr)
		}

		if ev == nil {
			panic("event is nil")
		}

		switch e := ev.(type) {
		case xproto.ClientMessageEvent:
			handler := env.get(e.Window)
			if handler.wmDeleteWindowAtom == xproto.Atom(e.Data.Data32[0]) {
				event.Type = backend.EventTypeClose
				event.Handler = handler
				return
			}

		case xproto.ExposeEvent:
			event.Type = backend.EventTypeExpose
			event.Handler = env.get(e.Window)
			event.Height = int(e.Height)
			event.Width = int(e.Width)
			return

		case xproto.MapNotifyEvent:
		case xproto.UnmapNotifyEvent:
		case xproto.ButtonPressEvent:
		case xproto.ButtonReleaseEvent:
		case xproto.MotionNotifyEvent:
		default:

		}
	}
}

func (env *Environment) put(windowID xproto.Window, handler *Handler) {
	env.mu.Lock()
	env.handlers[windowID] = handler
	env.mu.Unlock()
}

func (env *Environment) get(windowID xproto.Window) (handler *Handler) {
	env.mu.Lock()
	handler = env.handlers[windowID]
	env.mu.Unlock()
	return
}

func (env *Environment) remove(windowID xproto.Window) {
	env.mu.Lock()
	delete(env.handlers, windowID)
	env.mu.Unlock()
}
