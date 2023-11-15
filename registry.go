package libprovider

import (
	"errors"
	"sync"
)

var (
	ErrProviderNotFound   = errors.New("not found")
	ErrNotSupported       = errors.New("not supported")
	ErrProviderDuplicated = errors.New("provider duplicated")
)

func NewRegistry() *Registry {
	return &Registry{
		mu:      sync.RWMutex{},
		entries: make(map[string]NewProviderFunc),
	}
}

type Registry struct {
	mu      sync.RWMutex
	entries map[string]NewProviderFunc
}

func (r *Registry) Register(name string, builder NewProviderFunc) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.entries[name]; ok {
		return ErrProviderDuplicated
	}

	r.entries[name] = builder

	return nil
}

func (r *Registry) Lookup(name string) (NewProviderFunc, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	builder, ok := r.entries[name]
	if !ok {
		return nil, ErrProviderNotFound
	}

	return builder, nil
}

func (r *Registry) Walk(walkFunc func(string, NewProviderFunc)) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for name, builder := range r.entries {
		walkFunc(name, builder)
	}
}
