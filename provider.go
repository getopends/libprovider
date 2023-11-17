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
	Logger    Logger
	Variables Variables
	Secrets   Variables
}

type Variables map[string]Variable

func (e Variables) Get(key string) (string, error) {
	if v, ok := e[key]; ok {
		return v.Value, nil
	}

	return "", ErrNotFound
}

type Variable struct {
	Value string
}

// Info ...
type Info struct {
	Code      string                        `json:"code"`
	Name      string                        `json:"name"`
	Slug      string                        `json:"slug"`
	Version   string                        `json:"Version"`
	Author    string                        `json:"author"`
	Variables map[string]VariableDefinition `json:"variables"`
	URL       string                        `json:"url"`
	Icon      string                        `json:"ico"`
}

type VariableDefinition struct {
	Default     *string
	Required    bool
	Secret      bool
	Description string
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

func String(v string) *string {
	return &v
}
