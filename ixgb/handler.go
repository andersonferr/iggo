package ixgb

import (
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/andersonferr/iggo/backend"
)

// Handler handles the native window.
type Handler struct {
	drawable backend.Drawable

	title         string
	width, height int
	x, y          int

	env      *Environment
	windowID xproto.Window
	gcID     xproto.Gcontext

	wmDeleteWindowAtom xproto.Atom
}

// alloc allocate resources when the environment starts running.
func (handler *Handler) alloc() {
	env := handler.env
	wid, err := xproto.NewWindowId(env.conn)
	if err != nil {
		panic(err)
	}

	const valueMask uint32 = xproto.CwEventMask | xproto.CwBackPixel

	var valueList []uint32

	if valueMask&xproto.CwBackPixel != 0 {
		valueList = append(valueList, env.screen.WhitePixel)
	}

	if valueMask&xproto.CwEventMask != 0 {
		valueList = append(valueList,
			xproto.EventMaskStructureNotify|
				xproto.EventMaskExposure|
				xproto.EventMaskButtonPress|
				xproto.EventMaskButtonRelease|
				xproto.EventMaskButtonMotion|
				xproto.EventMaskPointerMotion|
				xproto.EventMaskKeyPress|
				xproto.EventMaskKeyRelease)
	}

	xproto.CreateWindow(
		env.conn,
		env.screen.RootDepth,
		wid,
		env.screen.Root,
		0, 0, // x, y
		uint16(handler.width), uint16(handler.height), // width, height
		2, // border
		xproto.WindowClassInputOutput,
		env.screen.RootVisual,
		valueMask, valueList)

	gcid, err := xproto.NewGcontextId(env.conn)
	if err != nil {
		panic(err)
	}
	xproto.CreateGC(env.conn, gcid, xproto.Drawable(wid), 0, nil)

	wmProtocolsAtom := getAtomOrPanic(env, "WM_PROTOCOLS")
	wmDeleteWindowAtom := getAtomOrPanic(env, "WM_DELETE_WINDOW")

	var data [4]byte
	xgb.Put32(data[:], uint32(wmDeleteWindowAtom))

	xproto.ChangeProperty(
		env.conn,
		xproto.PropModeReplace,
		wid,
		wmProtocolsAtom, xproto.AtomAtom,
		32, 1, data[:],
	)

	handler.gcID = gcid
	handler.windowID = wid
	handler.wmDeleteWindowAtom = wmDeleteWindowAtom
	handler.SetVisibility(true)
}

// SetVisibility of the native window.
func (handler *Handler) SetVisibility(visibility bool) {
	var err error
	if visibility {
		err = xproto.MapWindowChecked(handler.env.conn, handler.windowID).Check()
	} else {
		err = xproto.UnmapWindowChecked(handler.env.conn, handler.windowID).Check()
	}

	if err != nil {
		panic(err)
	}
}

// SetDrawable sets the drawable.
func (handler *Handler) SetDrawable(drawable backend.Drawable) {
	handler.drawable = drawable
}

// Drawable gets the drawable
func (handler *Handler) Drawable() backend.Drawable {
	return handler.drawable
}

func (handler *Handler) Deployer() backend.Deployer {
	return newDeployer(handler.env, handler.windowID, handler.gcID)
}

// Destroy the handler freeing the resources allocated.
func (handler *Handler) Destroy() {
	panic("not implemented")
}

// getAtomOrPanic gets the atom or panic
func getAtomOrPanic(env *Environment, atomName string) xproto.Atom {
	reply, err := xproto.InternAtom(
		env.conn,
		true,
		uint16(len(atomName)),
		atomName,
	).Reply()

	if err != nil {
		panic(err)
	}

	return reply.Atom
}

var (
	mu                 sync.Mutex
	mapWindowToHandler map[xproto.Window]*Handler
)

func init() {
	mapWindowToHandler = make(map[xproto.Window]*Handler)
}

func put(windowID xproto.Window, handler *Handler) {
	mu.Lock()
	mapWindowToHandler[windowID] = handler
	mu.Unlock()
}

func get(windowID xproto.Window) (handler *Handler) {
	mu.Lock()
	handler = mapWindowToHandler[windowID]
	mu.Unlock()
	return
}
