package ixgb

import "github.com/andersonferr/iggo/backend"

func init() {
	backend.Register(&XGBBackend{})
}

// XGBBackend is an implementation backend.Backend.
type XGBBackend struct{}

// Name returns the name of backend.
func (bk *XGBBackend) Name() string {
	return "xgb"
}

// Create creates a new environment, or returns an error.
func (bk *XGBBackend) Create() (backend.Environment, error) {
	return &Environment{}, nil
}
