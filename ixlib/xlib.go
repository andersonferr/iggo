package ixlib

/*
#cgo LDFLAGS: -lX11

#include <X11/Xlib.h>
#include <X11/Xutil.h>

#include <stdlib.h>
#include <stdint.h>

#include "IggoUtil.h"
*/
import "C"

import (
	"unsafe"
)

type (
	xDisplay  = *C.Display
	xWindow   = C.Window
	xPixmap   = C.Pixmap
	xEvent    = C.XEvent
	xGC       = C.GC
	xDrawable = C.Drawable
	xAtom     = C.Atom
	xStatus   = C.Status
	xID       = C.int
	xInt      = C.int

	xAnyEvent           = C.XAnyEvent
	xExposeEvent        = C.XExposeEvent
	xClientMessageEvent = C.XClientMessageEvent
	xConfigureEvent     = C.XConfigureEvent
	xBool               = C.Bool

	xCreateWindowStruct = C.IggoCreateWindowStruct
)

func xDestroyWindow(display xDisplay, window xWindow) {
	C.XDestroyWindow(display, window)
}

func xCreatePixmap(
	display xDisplay,
	d xDrawable,
	width int, height int,
	depth uint32,
) xPixmap {
	return C.XCreatePixmap(display, d, C.uint(width), C.uint(height), C.uint(depth))
}

func xCopyArea(
	display xDisplay,
	src, dest xDrawable,
	gc xGC,
	srcX, srcY int,
	width, height int,
	destX, destY int,
) {
	C.XCopyArea(
		display,
		src, dest,
		gc,
		xInt(srcX), xInt(srcY),
		C.uint(width), C.uint(height),
		xInt(destX), xInt(destY),
	)
}

func xMapWindow(display xDisplay, window xWindow) {
	C.XMapWindow(display, window)
}

func xUnmapWindow(display xDisplay, window xWindow) {
	C.XUnmapWindow(display, window)
}

func xNextEvent(d xDisplay, e *xEvent) {
	C.XNextEvent(d, e)
}

func xFlush(display xDisplay) {
	C.XFlush(display)
}

func xSync(display xDisplay, disgard xBool) {
	C.XSync(display, disgard)
}

func xDrawLine(display xDisplay, drawable xDrawable, gc xGC, x1, y1, x2, y2 int) {
	C.XDrawLine(
		display,
		drawable,
		gc,
		C.int(x1),
		C.int(y1),
		C.int(x2),
		C.int(y2),
	)
}

func xDrawPoint(display xDisplay, drawable xDrawable, gc xGC, x1, y1 int) {
	C.XDrawPoint(
		display,
		drawable,
		gc,
		C.int(x1),
		C.int(y1),
	)
}

func xDrawString(display xDisplay, drawable xDrawable, gc xGC, x1, y1 int, text string) {
	// cx := C.int(x)
	// cy := C.int(y) + C.int(d.font.ascent-d.font.descent)
	C.XDrawString(
		display,
		drawable,
		gc,
		C.int(x1),
		C.int(y1),
		C.CString(text),
		C.int(len(text)),
	)
}

func xOpenDisplay() xDisplay {
	return C.XOpenDisplay(nil)
}

func xPending(display xDisplay) int {
	return int(C.XPending(display))
}

func xDefaultScreen(display xDisplay) C.int {
	return C.XDefaultScreen(display)
}

func xDefaultRootWindow(display xDisplay) xWindow {
	return C.XDefaultRootWindow(display)
}

func xCloseDisplay(display xDisplay) {
	C.XCloseDisplay(display)
}

func xCreateSimpleWindow(display xDisplay, parent xWindow, x, y int, width, height uint) xWindow {
	screen := C.XDefaultScreen(display)
	white := C.XWhitePixel(display, screen)
	black := C.XBlackPixel(display, screen)
	root := C.XDefaultRootWindow(display)

	return C.XCreateSimpleWindow(
		display,        // display
		root,           // root window
		0,              // x position
		0,              // y position
		C.uint(width),  // width
		C.uint(height), // height
		0,              // border
		black,          // border pixel
		white,          // background
	)
}

const (
	xExposureMask        = C.ExposureMask
	xKeyPressMask        = C.KeyPressMask
	xButtonPressMask     = C.ButtonPressMask
	xStructureNotifyMask = C.StructureNotifyMask

	xExpose          = C.Expose
	xClientMessage   = C.ClientMessage
	xConfigureNotify = C.ConfigureNotify

	xTrue  = C.True
	xFalse = C.False
)

func xSelectInput(display xDisplay, window xWindow, mask uint64) {
	C.XSelectInput(display, window, C.long(mask))
}

func xCreateGC(display xDisplay, drawable xDrawable) xGC {
	return C.XCreateGC(display, drawable, 0, nil)
}

func xWhitePixel(display xDisplay, screen C.int) uint64 {
	return uint64(C.XWhitePixel(display, screen))
}
func xBlackPixel(display xDisplay, screen C.int) uint64 {
	return uint64(C.XBlackPixel(display, screen))
}

func xSetBackground(display xDisplay, gc xGC, color uint64) {
	C.XSetBackground(display, gc, C.ulong(color))
}

func xSetForeground(display xDisplay, gc xGC, color uint64) {
	C.XSetForeground(display, gc, C.ulong(color))
}

func xInternAtom(display xDisplay, atomName string, onlyIfExist bool) xAtom {
	return C.XInternAtom(display, C.CString(atomName), map[bool]C.Bool{true: C.True, false: C.False}[onlyIfExist])
}

func xSetWMProtocols(display xDisplay, window xWindow, protocols []xAtom) xStatus {
	return C.XSetWMProtocols(display, window, &protocols[0], C.int(len(protocols)))
}

// Font -----------------------------------------------------------------------
type (
	xFont   = *C.XFontStruct
	xFontID = C.Font
)

func xLoadQueryFont(display xDisplay, name string) xFont {
	return C.XLoadQueryFont(display, C.CString(name))
}

func xTextWidth(font xFont, text string) int {
	return int(C.XTextWidth(font, C.CString(text), xInt(len(text))))
}

func xSetFont(display xDisplay, gc xGC, fontID xFontID) {
	C.XSetFont(display, gc, fontID)
}

func readEvents(display xDisplay, buffer []xEvent) int {
	return int(C.IggoReadEvents(display, (*xEvent)(unsafe.Pointer(&buffer[0])), xInt(len(buffer))))
}

func drawImage(display xDisplay, drawable xDrawable, gc xGC, imagedata []byte, width, height int, bytesPerLine int) {
	C.IggoDrawImage(
		display,
		drawable,
		gc,
		(*C.uchar)(unsafe.Pointer(&imagedata[0])),
		xInt(width),
		xInt(height),
		xInt(bytesPerLine),
	)
}

func xCreateWindow(display xDisplay, width, height int) xCreateWindowStruct {
	r := C.IggoCreateWindow(
		display,
		xInt(width),
		xInt(height),
	)

	return r
}
