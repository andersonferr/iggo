package backend_test

import (
	"testing"

	"github.com/andersonferr/iggo/backend"
)

func TestEnvironmentManagerFuncRegisterProviderWithEmptyName(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("EnvironmentManager.RegisterProvider: expected an error!")
		}
	}()

	em := backend.NewEnvironmentManager()
	em.RegisterProvider("", nil)
}
