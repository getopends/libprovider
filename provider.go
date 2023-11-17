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
type Options struct{}

// Info ...
type Info struct {
	Code    string
	Name    string
	Slug    string
	Version string
	Author  string
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
