package libprovider

import (
	"sync"

	"github.com/getopends/libprovider"
)

// New
func New() *Registry {
	return &Registry{
		l:       sync.RWMutex{},
		entries: make(map[string]libprovider.NewProviderFunc),
	}
}

// Registry
type Registry struct {
	l       sync.RWMutex
	entries map[string]libprovider.NewProviderFunc
}

// Register
func (r *Registry) Register(name string, builder libprovider.NewProviderFunc) error {
	r.l.Lock()
	defer r.l.Unlock()

	if _, ok := r.entries[name]; ok {
		return libprovider.ErrProviderDuplicated
	}

	r.entries[name] = builder

	return nil
}

// Lookup
func (r *Registry) Lookup(name string) (libprovider.NewProviderFunc, error) {
	r.l.RLock()
	defer r.l.RUnlock()

	builder, ok := r.entries[name]
	if !ok {
		return nil, libprovider.ErrProviderNotFound
	}

	return builder, nil
}

// Walk
func (r *Registry) Walk(walkFunc func(string, libprovider.NewProviderFunc)) {
	r.l.RLock()
	defer r.l.RUnlock()

	for name, builder := range r.entries {
		walkFunc(name, builder)
	}
}
