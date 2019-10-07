package ixgb

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/andersonferr/iggo/backend"
)

// BackendName the name of backend
const BackendName = "XGB"

// Environment implements backend.Environment for XGB.
type Environment struct {
	mu sync.Mutex

	handlers []Handler

	conn   *xgb.Conn
	screen *xproto.ScreenInfo
}

func init() {
	// help typechecker
	var env *Environment
	_ = backend.Environment(env)
}

func init() {
	backend.Register(BackendName, func() backend.Environment {
		return &Environment{}
	})
}

func (env *Environment) Run() error {
	if len(env.handlers) == 0 {
		return nil
	}

	X, err := xgb.NewConn()
	if err != nil {
		return err
	}

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)

	env.conn = X
	env.screen = screen
	defer env.conn.Close()

	for i := 0; i < len(env.handlers); i++ {
		env.handlers[i].alloc()
		fmt.Printf("Handler: %#v\n", env.handlers[i])
	}

	for len(env.handlers) != 0 {
		ev, xerr := env.conn.WaitForEvent()
		if ev == nil && xerr == nil {
			panic("nil answer")
		}

		if ev != nil {
			switch e := ev.(type) {
			case xproto.ClientMessageEvent:
				handler := env.findHandler(e.Window)
				if e.Data.Data32[0] == uint32(handler.wmDeleteWindowAtom) {
					err := xproto.DestroyWindowChecked(env.conn, e.Window).Check()
					if err != nil {
						panic(err)
					}
					env.removeHandler(handler)
				}

			case xproto.ExposeEvent:
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
			return (xerr)
		}
	}

	return nil
}

// CreateHandler creates a new handler.
func (env *Environment) CreateHandler(title string, x, y, width, height int) (backend.Handler, error) {
	handler := Handler{
		env:    env,
		title:  title,
		width:  width,
		height: height,
		x:      x,
		y:      y,
	}

	env.handlers = append(env.handlers, handler)

	return &handler, nil
}

func (env *Environment) findHandler(wid xproto.Window) *Handler {
	for i := 0; i < len(env.handlers); i++ {
		if env.handlers[i].windowID == wid {
			return &env.handlers[i]
		}
	}

	return nil
}

func (env *Environment) removeHandler(handler *Handler) {
	for i := 0; i < len(env.handlers); i++ {
		if &env.handlers[i] == handler {
			env.handlers = append(env.handlers[:i], env.handlers[i+1:]...)
		}
	}
}
