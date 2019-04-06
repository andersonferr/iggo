package ixlib

import (
	"image"

	"github.com/andersonferr/iggo/backend"
)

type Handler struct {
	backend.BaseHandler

	display xDisplay
	window  xWindow
	gc      xGC
	fb      *FrameBuffer
}

func (env *Environment) CreateHandler(width, height int) backend.Handler {
	if !env.started {
		panic(errEnvUninitieded)
	}

	r := xCreateWindow(env.display, width, height)

	h := &Handler{
		display: env.display,
		window:  r.window,
		gc:      r.gc,
		fb:      NewFrameBuffer(image.Rect(0, 0, 2000, 2000)),
	}
	xWindow2Handler[r.window] = h

	return h
}

func (h *Handler) SetVisibility(visibility bool) {
	if visibility {
		xMapWindow(h.display, h.window)
	} else {
		xUnmapWindow(h.display, h.window)
	}
	xFlush(h.display)
}

func (h *Handler) Destroy() {
	xDestroyWindow(h.display, h.window)
}

func (h *Handler) Deployer() backend.Deployer {
	return h
}

func (h *Handler) Deploy(im *image.RGBA, area image.Rectangle) {
	area = area.Intersect(h.fb.Rect)
	// fast reject
	if area.Empty() {
		return
	}

	for x := area.Min.X; x < area.Max.X; x++ {
		for y := area.Min.Y; y < area.Max.Y; y++ {
			h.fb.SetRGBA(x, y, im.RGBAAt(x, y))
		}
	}

	drawImage(
		h.display,
		xDrawable(h.window),
		h.gc,
		h.fb.Pix,
		area.Dx(),
		area.Dy(),
		h.fb.Stride,
	)
}
