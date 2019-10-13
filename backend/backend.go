package backend

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Backend creates a new environment.
type Backend interface {
	// Name is the name of backend. the name must be constant.
	Name() string

	// Create creates a new environment.
	Create() (Environment, error)
}

var (
	backends   = make(map[string]Backend)
	backendsMu sync.RWMutex
)

// Get gets a registered backend by name, or returns an error;
func Get(name string) (Backend, error) {
	backendsMu.RLock()
	bk, ok := backends[name]
	backendsMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("backend %q is not registered (forgot to import?)", name)
	}

	return bk, nil
}

// Register registers a backend.
func Register(backend Backend) {
	if backend == nil {
		panic(errors.New("backend must not be nil"))
	}

	name := backend.Name()
	if name == "" {
		panic(errors.New("backend must have a name"))
	}

	backendsMu.Lock()
	backends[name] = backend
	backendsMu.Unlock()
}

// List returns a list of all registered backends.
func List() []string {
	backendsMu.RLock()
	defer backendsMu.RUnlock()

	l := make([]string, len(backends), len(backends))
	i := 0
	for k := range backends {
		l[i] = k
		i++
	}

	sort.Strings(l)

	return l
}
