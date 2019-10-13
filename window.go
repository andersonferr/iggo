package iggo

import (
	"image"

	"github.com/andersonferr/iggo/backend"
)

//Window holds native-window handler
type Window struct {
	// handler is a handler of native window.
	handler backend.Handler

	// buffer is the buffer where the image is drawn.
	buffer *image.RGBA

	container BasicContainer

	resizeCallback func(window *Window, width, height int)
	closeCallback  func(window *Window)

	// title
	title string

	// position
	x, y int

	// dimensions
	width, height int
}

func init() {
	var w *Window
	_ = Container(w)
}

func CreateWindow(title string, x, y, width, height int) *Window {
	if width < 0 {
		width = 0
	}

	if height < 0 {
		height = 0
	}

	return &Window{
		title:  title,
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (w *Window) SetVisibility(visibility bool) {
	w.handler.SetVisibility(visibility)
}

// resize deve ser usada para atualizar o tamanho da janela virtual(window)
// e chama o callback de redimensionamento(resizeCallback) para a janela em questão.
// esta função não redimensiona a janela nativa(handler) e deve ser chamada apenas quando
// a janela nativa(handler) for redimensionada.
func resize(window *Window, width, height int) {
	if width < 0 {
		width = 0
	}
	if height < 0 {
		height = 0
	}

	window.container.width = width
	window.container.height = height

	size := width * height * 4 // espaço extra
	if len(window.buffer.Pix) < size {
		window.buffer = image.NewRGBA(image.Rect(0, 0, width, height))
	}

	if callback := window.resizeCallback; callback != nil {
		callback(window, width, height)
	}
}

// OnResize add resize event listener or remove if callback is nil.
func (w *Window) OnResize(callback func(window *Window, width, height int)) {
	w.resizeCallback = callback
}

// OnClose add close event listener or remove if callback is nil.
func (w *Window) OnClose(callback func(window *Window)) {
	w.closeCallback = callback
}

func (w *Window) Close() {
	w.handler.Destroy()
	w.handler = nil
}

func (w *Window) Parent() Container {
	return nil
}

func (w *Window) Height() int {
	return w.container.Height()
}

func (w *Window) SetHeight(height int) {
	panic("not implemented")
}

func (w *Window) Width() int {
	return w.container.Width()
}

func (w *Window) SetWidth(width int) {
	panic("not implemented")
}

// Draw must not be called!
// it panics.
func (w *Window) Draw(drawer Drawer) {
	panic("Draw must not be called!")
}

func (w *Window) Add(widget Widget) {
	w.container.Add(widget)
}

func (w *Window) Remove(widget Widget) {
	w.container.Remove(widget)
}

func (w *Window) Children() []Widget {
	return w.container.Children()
}
