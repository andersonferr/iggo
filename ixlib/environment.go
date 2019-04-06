package ixlib

import (
	"errors"

	"github.com/andersonferr/iggo/backend"
)

var (
	errEnvUninitieded = errors.New("environment unitialized")
)

type Environment struct {
	started bool
	display xDisplay
	screen  xID
	root    xWindow
}

func init() {
	// help type checker
	var e *Environment
	_ = backend.Environment(e)
}

func (env *Environment) Start() {
	display := xOpenDisplay()
	if display == nil {
		panic(errors.New("could not open x display"))
	}
	scr := xDefaultScreen(display)
	root := xDefaultRootWindow(display)

	env.screen = scr
	env.display = display
	env.root = root
	env.started = true
}

func (env *Environment) Finish() {
	if !env.started {
		return
	}

	xCloseDisplay(env.display)
	env.started = false
}
