package libprovider

import (
	"sync"

	"github.com/getopends/libprovider"
)

// New
func New() *Registry {
	return &Registry{
		l:       sync.RWMutex{},
		entries: make(map[string]libprovider.BuildFunc),
	}
}

// Registry
type Registry struct {
	l       sync.RWMutex
	entries map[string]libprovider.BuildFunc
}

// Register
func (r *Registry) Register(name string, builder libprovider.BuildFunc) error {
	r.l.Lock()
	defer r.l.Unlock()

	if _, ok := r.entries[name]; ok {
		return libprovider.ErrDuplicated
	}

	r.entries[name] = builder

	return nil
}

// Lookup
func (r *Registry) Lookup(name string) (libprovider.BuildFunc, error) {
	r.l.RLock()
	defer r.l.RUnlock()

	builder, ok := r.entries[name]
	if !ok {
		return nil, libprovider.ErrNotFound
	}

	return builder, nil
}

// Walk
func (r *Registry) Walk(walkFunc func(string, libprovider.BuildFunc)) {
	r.l.RLock()
	defer r.l.RUnlock()

	for name, builder := range r.entries {
		walkFunc(name, builder)
	}
}
