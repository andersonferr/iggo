package backend

import (
	"image"
	"sync"
)

//Deployer é responsável por instalar a imagem na tela(screen)
type Deployer interface {
	//Deploy copia a imagem im para a tela(screen). a imagem im deve ter ordem MSB(Most Significant Byte First)
	Deploy(im *image.RGBA, area image.Rectangle)
}

// EnvironmentProvider provides a new environment.
type EnvironmentProvider func() Environment

var (
	handlers = map[string]EnvironmentProvider{}
	mutex    sync.Mutex
)

func Get(name string) EnvironmentProvider {
	mutex.Lock()
	defer mutex.Unlock()

	return handlers[name]
}

func Register(name string, provider EnvironmentProvider) {
	mutex.Lock()
	defer mutex.Unlock()

	handlers[name] = provider
}
