package iggo

type Drawer interface {
	DrawRect(x0, y0, x1, y1 float64)
	DrawText(text string, x0, y0, x1, y1 float64)
	Fill()
	Stroke()
	Clear()
	SetColor(r, g, b, a uint8)
}
