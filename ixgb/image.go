package ixgb

import (
	"image"
	"image/color"
	"image/draw"
)

type argbImage struct {
	pix  []uint8
	rect image.Rectangle
}

func init() {
	var im *argbImage
	_ = draw.Image(im)
}

func newArgbImage(r image.Rectangle) *argbImage {
	im := &argbImage{}
	im.SetBounds(r)
	return im
}

func (im *argbImage) SetBounds(r image.Rectangle) {
	size := r.Dx() * r.Dy() * 4
	if len(im.pix) < size {
		im.pix = make([]uint8, size)
	}

	im.rect = r
}

func (im *argbImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (im *argbImage) Bounds() image.Rectangle {
	return im.rect
}

func (im *argbImage) At(x, y int) color.Color {
	const (
		a = 0
		r = 1
		g = 2
		b = 3
	)

	if !(image.Point{x, y}.In(im.rect)) {
		return color.RGBA{}
	}

	i := (y*im.rect.Dx() + x) * 4
	s := im.pix[i : i+4 : i+4]

	return color.RGBA{R: s[r], G: s[g], B: s[b], A: s[a]}
}

func (im *argbImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(im.rect)) {
		return
	}
	c0 := color.RGBAModel.Convert(c).(color.RGBA)
	i := (y*im.rect.Dx() + x) * 4
	im.pix[i+0] = c0.A
	im.pix[i+1] = c0.R
	im.pix[i+2] = c0.G
	im.pix[i+3] = c0.B
}

func (im *argbImage) SetRGBA(x, y int, c color.RGBA) {
	if !(image.Point{x, y}.In(im.rect)) {
		return
	}
	i := (y*im.rect.Dx() + x) * 4
	im.pix[i+0] = c.A
	im.pix[i+1] = c.R
	im.pix[i+2] = c.G
	im.pix[i+3] = c.B
}
