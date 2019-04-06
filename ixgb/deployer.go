package ixgb

import (
	"image"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/andersonferr/iggo/backend"
)

type Deployer struct {
	env      *Environment
	windowID xproto.Window
	gcID     xproto.Gcontext

	im *argbImage
}

func newDeployer(env *Environment, wid xproto.Window, gcid xproto.Gcontext) backend.Deployer {
	w := env.screen.WidthInPixels
	h := env.screen.HeightInPixels

	return &Deployer{
		env:      env,
		windowID: wid,
		gcID:     gcid,
		im:       newArgbImage(image.Rect(0, 0, int(w), int(h))),
	}
}

func (deployer *Deployer) Deploy(im *image.RGBA, area image.Rectangle) {
	area = area.Intersect(deployer.im.rect)
	// fast reject
	if area.Empty() {
		return
	}

	for x := area.Min.X; x < area.Max.X; x++ {
		for y := area.Min.Y; y < area.Max.Y; y++ {
			deployer.im.SetRGBA(x, y, im.RGBAAt(x, y))
		}
	}

	width := area.Dx()
	height := area.Dy()

	xproto.PutImage(
		deployer.env.conn,
		xproto.ImageFormatZPixmap,
		xproto.Drawable(deployer.windowID),
		deployer.gcID,
		uint16(width),
		uint16(height),
		0, 0,
		0,
		deployer.env.screen.RootDepth,
		deployer.im.pix[:width*height*4])
}
