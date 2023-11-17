package libprovider

import (
	"context"
	"errors"
)

var (
	// ErrNotFound ...
	ErrNotFound = errors.New("not found")

	// ErrNotSupported ...
	ErrNotSupported = errors.New("not supported")

	// ErrDuplicated ...
	ErrDuplicated = errors.New("provider duplicated")
)

// Options ...
type Options struct {
	Logger  Logger
	Env     Envs
	Secrets Envs
}

type Envs map[string]Env

func (e Envs) Get(key string) (string, error) {
	if v, ok := e[key]; ok {
		return v.Value, nil
	}

	return "", ErrNotFound
}

type Env struct {
	Value string
}

// Info ...
type Info struct {
	Code    string   `json:"code"`
	Name    string   `json:"name"`
	Slug    string   `json:"slug"`
	Version string   `json:"Version"`
	Author  string   `json:"author"`
	Env     []string `json:"variables"`
	Secrets []string `json:"secrets"`
}

// BuildFunc ...
type BuildFunc func(context.Context) Provider

// Provider ...
type Provider interface {
	Info() Info
	Configure(context.Context, *Options) error
}

// Error ...
type Error struct{}

func (e Error) Error() string {
	return "error"
}

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}
