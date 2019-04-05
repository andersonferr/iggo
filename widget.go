package iggo

//Widget is a component of GUI.
type Widget interface {
	Parent() Container

	Height() int
	SetHeight(height int)

	Width() int
	SetWidth(width int)

	Draw(drawer Drawer)
}

//Container is a widget that can hold others widgets.
type Container interface {
	Widget

	Add(widget Widget)
	Remove(widget Widget)

	Children() []Widget
}

// BasicWidget is an basic implementation of Widget.
type BasicWidget struct {
	width, height int //dimention
	parent        Container
}

//Parent returns the box where container is inside.
func (widget *BasicWidget) Parent() Container {
	return widget.parent
}

//Draw the container into screen
func (widget *BasicWidget) Draw(drawer Drawer) {
	// do nothing
}

//Height returns the height of container
func (widget *BasicWidget) Height() int {
	return widget.height
}

//SetHeight set the height of container
func (widget *BasicWidget) SetHeight(height int) {
	height = max(height, 0)
	if widget.height == height {
		return
	}
	widget.height = height
	// widget.needRelayout = true
}

//Width returns the width of container
func (widget *BasicWidget) Width() int {
	return widget.width
}

//SetWidth set the width of container
func (widget *BasicWidget) SetWidth(width int) {
	width = max(width, 0)
	if widget.width == width {
		return
	}

	widget.width = width
}

// BasicContainer is an basic implementation of Container.
type BasicContainer struct {
	BasicWidget
	children []Widget

	needRelayout bool
}

//Children returns the children.
func (container *BasicContainer) Children() []Widget {
	return container.children
}

//Add widget as child.
func (container *BasicContainer) Add(widget Widget) {
	if widget == nil {
		return
	}

	parent := widget.Parent()
	if parent != nil {
		parent.Remove(widget)
	}

	container.children = append(container.children, widget)
}

//Remove the given box from container if has it
func (container *BasicContainer) Remove(widget Widget) {
	if widget == nil {
		return
	}

	for i := 0; i < len(container.children); i++ {
		if container.children[i] == widget {
			container.children = append(container.children[0:i], container.children[i+1:]...)
			break
		}
	}
}
