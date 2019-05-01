package backend

import (
	"image"
)

type Handler interface {
	SetVisibility(visibility bool)
	Deployer() Deployer
	Destroy()

	Data() interface{}
	SetData(interface{})
}

//Deployer é responsável por instalar a imagem na tela(screen)
type Deployer interface {
	//Deploy copia a imagem im para a tela(screen). a imagem im deve ter ordem MSB(Most Significant Byte First)
	Deploy(im *image.RGBA, area image.Rectangle)
}

// Environment is where the gui will run.
type Environment interface {
	// CreateHandler creates a new handler (native window) for a iggo window.
	CreateHandler(width, height int) Handler

	// NextEvent gets next event.
	NextEvent(*Event)

	// Start prepare for listen events.
	Start()

	//Finish clean the environment after listen events.
	Finish()
}

type BaseHandler struct {
	data interface{}
}

func (bh *BaseHandler) Data() interface{} {
	return bh.data
}

func (bh *BaseHandler) SetData(data interface{}) {
	bh.data = data
}
