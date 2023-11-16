package libprovider

import (
	"sync"
)

// NewRegistry
func NewRegistry() *Registry {
	return &Registry{
		mu:      sync.RWMutex{},
		entries: make(map[string]NewProviderFunc),
	}
}

// Registry
type Registry struct {
	mu      sync.RWMutex
	entries map[string]NewProviderFunc
}

// Register
func (r *Registry) Register(name string, builder NewProviderFunc) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.entries[name]; ok {
		return ErrProviderDuplicated
	}

	r.entries[name] = builder

	return nil
}

// Lookup
func (r *Registry) Lookup(name string) (NewProviderFunc, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	builder, ok := r.entries[name]
	if !ok {
		return nil, ErrProviderNotFound
	}

	return builder, nil
}

// Walk
func (r *Registry) Walk(walkFunc func(string, NewProviderFunc)) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for name, builder := range r.entries {
		walkFunc(name, builder)
	}
}
