package iggo

import (
	"errors"
	"sync"

	"github.com/andersonferr/iggo/backend"
)

// Application holds the windows
type Application struct {
	env backend.Environment

	windows []*Window
	mu      sync.Mutex
}

// NewApplication creates a new application with default environment as backend.
func NewApplication() *Application {
	return NewApplicationWithEnvironment(env)
}

// NewApplicationWithEnvironment creates a new application with given environment as backend.
func NewApplicationWithEnvironment(env backend.Environment) *Application {
	return &Application{
		env: env,
	}
}

// Run the application.
func (app *Application) Run() error {
	if app == nil {
		return errors.New("nil application")
	}

	env := app.env
	if env == nil {
		return errors.New("nil environment")
	}

	return env.Run()
}

func (app *Application) AddWindow(w *Window) {
	handler, err := app.env.CreateHandler(w.title, w.x, w.y, w.width, w.height)

	if err != nil {
		panic(err)
	}

	handler.SetDrawable(w)
	w.handler = handler
}
