package ixlib

import "github.com/andersonferr/iggo/backend"

func init() {
	backend.Register("X11", func() backend.Environment {
		return &Environment{}
	})
}
