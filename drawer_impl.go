package iggo

import (
	"image"

	"github.com/fogleman/gg"
)

type idrawer struct {
	ctx *gg.Context
}

func init() {
	var d *idrawer
	_ = Drawer(d)
}

func NewDrawer(im *image.RGBA) Drawer {
	ggctx := gg.NewContextForRGBA(im)
	return &idrawer{
		ctx: ggctx,
	}
}

func (d *idrawer) DrawRect(x0, y0, x1, y1 float64) {
	d.ctx.DrawRectangle(x0, y0, x1-x0, y1-y0)
}

func (d *idrawer) DrawText(text string, x0, y0, x1, y1 float64) {
	d.ctx.DrawString(text, x0, y0)
}

func (d *idrawer) Fill() {
	d.ctx.Fill()
}
func (d *idrawer) Stroke() {
	d.ctx.Stroke()
}
func (d *idrawer) Clear() {
	d.ctx.Clear()
}

func (d *idrawer) SetColor(r, g, b, a uint8) {
	d.ctx.SetRGBA255(int(r), int(g), int(b), int(a))
}
