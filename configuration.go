package libprovider

import (
	"context"
)

// Configuration
type Configuration interface {
	Configure(context.Context, *Options) error
}

// ConfigurationFunc
type ConfigurationFunc func(context.Context) error

// Configure
func (c ConfigurationFunc) Configure(ctx context.Context) error {
	return c(ctx)
}
