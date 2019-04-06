package ixgb

import (
	"fmt"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/andersonferr/iggo/backend"
)

// Environment implements backend.Environment for XGB.
type Environment struct {
	conn   *xgb.Conn
	screen *xproto.ScreenInfo
}

func init() {
	// help typechecker
	var env *Environment
	_ = backend.Environment(env)
}

func init() {
	backend.Register("XGB", func() backend.Environment {
		return &Environment{}
	})
}

// CreateHandler create a new backend.Handler for this environment.
func (env *Environment) CreateHandler(width, height int) backend.Handler {
	return newHandler(env, width, height)
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
	ev, xerr := env.conn.WaitForEvent()
	if ev == nil && xerr == nil {
		panic(xerr)
	}

	if ev != nil {
		switch e := ev.(type) {
		case xproto.ClientMessageEvent:
			event.Type = backend.EventTypeClose
			event.Handler = get(e.Window)

		case xproto.ExposeEvent:
			event.Type = backend.EventTypeExpose
			event.Height = int(e.Height)
			event.Width = int(e.Width)

		case xproto.MapNotifyEvent:
		case xproto.UnmapNotifyEvent:
		case xproto.ButtonPressEvent:
		case xproto.ButtonReleaseEvent:
		case xproto.MotionNotifyEvent:
		default:
		}
	}
	if xerr != nil {
		// fmt.Printf("Error: %s\n", xerr)
		panic(xerr)
	}
}
