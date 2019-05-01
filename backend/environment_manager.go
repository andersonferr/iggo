package backend

import (
	"sync"
)

//EnvironmentManager manages the environments.
type EnvironmentManager struct {
	mu sync.Mutex
	ep map[string]EnvironmentProvider
}

// NewEnvironmentManager creates a new EnvironmentManager intance.
func NewEnvironmentManager() *EnvironmentManager {
	em := &EnvironmentManager{
		ep: make(map[string]EnvironmentProvider),
	}

	return em
}

// EnvironmentProvider provides a new environment.
type EnvironmentProvider func() Environment

//DefaultEnvironmentManager is the default environment.
var DefaultEnvironmentManager = NewEnvironmentManager()

//RegisterProvider register a new environment if ep is not nil, else, remove it.
//panics if name is "".
func (em *EnvironmentManager) RegisterProvider(name string, ep EnvironmentProvider) {
	if name == "" {
		panic("EnvironmentManager.RegisterProvider: name must be set.")
	}

	em.mu.Lock()
	defer em.mu.Unlock()

	if ep != nil {
		em.ep[name] = ep
	} else {
		delete(em.ep, name)
	}
}

//GetProvider returns a registered provider, or nil if name is not registered.
//panics if name is "".
func (em *EnvironmentManager) GetProvider(name string) EnvironmentProvider {

	if name == "" {
		panic("EnvironmentManager.GetProvider: name must be set.")
	}

	em.mu.Lock()
	defer em.mu.Unlock()
	return em.ep[name]
}

//CreateEnvironment creates a new environment based on registered provider name.
//returns nil if provider names not registered.
func (em *EnvironmentManager) CreateEnvironment(name string) Environment {
	ep := em.GetProvider(name)
	if ep != nil {
		return ep()
	}

	return nil
}

// List returns the list of registered environments.
func (em *EnvironmentManager) List() []string {
	em.mu.Lock()
	defer em.mu.Unlock()

	l := []string{}
	for k := range em.ep {
		l = append(l, k)
	}

	return l
}
