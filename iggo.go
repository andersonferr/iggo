package iggo

import (
	"fmt"
	"runtime"

	"github.com/andersonferr/iggo/backend"
)

var env backend.Environment

//Use set the environment provider to be used when create a new environment.
func Use(name string) error {
	provider := backend.Get(name)
	if provider == nil {
		return fmt.Errorf("backend %q is not registered", name)
	}

	env = provider()
	return nil
}

//Run runs the application main function after prepare the environment for that
func Run(fn func()) {
	runtime.LockOSThread()

	closeRequested = false
	env.Start()
	defer env.Finish()

	fn()
	handleEvents()
}

var (
	handlerToWindow = map[backend.Handler]*Window{}
	closeRequested  bool
)

func handleEvents() {
	var e backend.Event
	for !closeRequested {
		env.NextEvent(&e)
		w := handlerToWindow[e.Handler]

		switch e.Type {
		case backend.EventTypeClose:
			if callback := w.closeCallback; callback != nil {
				callback(w)
			}
			w.Close()

		case backend.EventTypeResize:
			if e.Width != w.Width() || e.Height != w.Height() {
				resize(w, e.Width, e.Height)
			}
		}

		for _, w := range handlerToWindow {
			w.draw()

		}

		//time.Sleep(30 * time.Millisecond)
	}
}

func Close() {
	closeRequested = true
}
