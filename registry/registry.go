package libprovider

import (
	"sync"

	"github.com/getopends/libprovider"
)

// New
func New() *Registry {
	return &Registry{
		mu:      sync.RWMutex{},
		entries: make(map[string]libprovider.NewProviderFunc),
	}
}

// Registry
type Registry struct {
	mu      sync.RWMutex
	entries map[string]libprovider.NewProviderFunc
}

// Register
func (r *Registry) Register(name string, builder libprovider.NewProviderFunc) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.entries[name]; ok {
		return libprovider.ErrProviderDuplicated
	}

	r.entries[name] = builder

	return nil
}

// Lookup
func (r *Registry) Lookup(name string) (libprovider.NewProviderFunc, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	builder, ok := r.entries[name]
	if !ok {
		return nil, libprovider.ErrProviderNotFound
	}

	return builder, nil
}

// Walk
func (r *Registry) Walk(walkFunc func(string, libprovider.NewProviderFunc)) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for name, builder := range r.entries {
		walkFunc(name, builder)
	}
}
