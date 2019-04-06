package ixgb

import (
	"image"
	"image/color"
	"image/draw"
)

type FrameBuffer struct {
	Pix    []uint8
	Rect   image.Rectangle
	Stride int
}

func init() {
	var fb *FrameBuffer
	_ = draw.Image(fb)
}

func NewFrameBuffer(r image.Rectangle) *FrameBuffer {
	w, h := r.Dx(), r.Dy()
	p := make([]uint8, w*h*4)

	return &FrameBuffer{
		Pix:    p,
		Rect:   r,
		Stride: w * 4,
	}
}

func (fb *FrameBuffer) ColorModel() color.Model {
	return color.RGBAModel
}

func (fb *FrameBuffer) Bounds() image.Rectangle {
	return fb.Rect
}

func (fb *FrameBuffer) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(fb.Rect)) {
		return color.RGBA{}
	}
	var r, g, b, a uint8
	i := y*fb.Stride + x*4
	a = fb.Pix[i+0]
	r = fb.Pix[i+1]
	g = fb.Pix[i+2]
	b = fb.Pix[i+3]

	return color.RGBA{r, g, b, a}
}

func (fb *FrameBuffer) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(fb.Rect)) {
		return
	}
	c0 := color.RGBAModel.Convert(c).(color.RGBA)
	i := y*fb.Stride + x*4
	fb.Pix[i+0] = c0.A
	fb.Pix[i+1] = c0.R
	fb.Pix[i+2] = c0.G
	fb.Pix[i+3] = c0.B
}

func (fb *FrameBuffer) SetRGBA(x, y int, c color.RGBA) {
	if !(image.Point{x, y}.In(fb.Rect)) {
		return
	}
	i := y*fb.Stride + x*4
	fb.Pix[i+0] = c.A
	fb.Pix[i+1] = c.R
	fb.Pix[i+2] = c.G
	fb.Pix[i+3] = c.B
}

func index() int {
	return 0
}
