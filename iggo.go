package iggo

import (
	"fmt"

	"github.com/andersonferr/iggo/backend"
)

var env backend.Environment

//Use set the environment provider to be used when create a new environment.
func Use(name string) error {
	provider := backend.Get(name)
	if provider == nil {
		return fmt.Errorf("backend %q is not registered", name)
	}

	env = provider()
	return nil
}
