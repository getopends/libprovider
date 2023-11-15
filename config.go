package libprovider

import (
	"context"
)

type Configurer interface {
	Configure(context.Context, *Options) error
}

type ConfigurerFunc func(context.Context) error

func (c ConfigurerFunc) Configure(ctx context.Context) error {
	return c(ctx)
}
