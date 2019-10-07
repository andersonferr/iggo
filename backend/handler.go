package backend

// Handler is the native window handler.
type Handler interface {
	SetDrawable(Drawable)
	Drawable() Drawable

	SetVisibility(bool)
	Destroy()
}
