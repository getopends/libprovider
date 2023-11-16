package libprovider

import (
	"context"
)

type Configuration interface {
	Configure(context.Context, *ConfigureOptions) error
}

type ConfigurationFunc func(context.Context) error

func (c ConfigurationFunc) Configure(ctx context.Context) error {
	return c(ctx)
}
