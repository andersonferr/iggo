package iggo

import (
	"log"
	"runtime"

	"github.com/andersonferr/iggo/backend"

	// Known backends
	_ "github.com/andersonferr/iggo/ixgb"
)

// DefaultBackend is the defalt backend.
var DefaultBackend backend.Backend

func init() {
	switch runtime.GOOS {
	case "linux":
		for _, name := range []string{"xgb", "xlib", "wayland"} {
			if bk, err := backend.Get(name); err == nil {
				DefaultBackend = bk
				return
			}
		}

	default:
	}

	log.Println("iggo do not support by default")
}
