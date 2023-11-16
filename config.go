package libprovider

import (
	"context"
)

type ProviderConfiguration interface {
	Configure(context.Context, *ConfigureOptions) error
}

type ConfigurerFunc func(context.Context) error

func (c ConfigurerFunc) Configure(ctx context.Context) error {
	return c(ctx)
}
