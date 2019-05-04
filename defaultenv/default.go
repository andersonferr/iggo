package defaultenv

import (
	"github.com/andersonferr/iggo"
	"github.com/andersonferr/iggo/backend"
)

func init() {
	provider := backend.DefaultEnvironmentManager.GetProvider(DefaultEnvName)
	if provider == nil {
		panic("defaultenv: default package ixgb is not available.")
	}

	backend.DefaultEnvironmentManager.RegisterProvider("default", provider)
	iggo.Use("default")
}
