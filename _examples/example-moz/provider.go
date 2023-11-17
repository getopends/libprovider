package example

import (
	"context"

	"github.com/getopends/libprovider"
)

var _ libprovider.Provider = &ExampleProvider{}

func NewExampleProvider(context.Context) libprovider.Provider {
	return &ExampleProvider{}
}

type ExampleProvider struct{}

func (p *ExampleProvider) Info() libprovider.Info {
	return libprovider.Info{
		Name:    "Example",
		Slug:    "example-moz",
		Author:  "Edson Michaque",
		Version: "0.1.0",
		Secrets: []string{
			"USERNAME",
			"PASSWORD",
			"ENCRYPTION_KEY"
		},
		Env: []string{
			"HOST",
			"PORT"
		},
	}
}

func (p *ExampleProvider) Configure(ctx context.Context, opts *libprovider.Options) error {
	return nil
}

func (p *ExampleProvider) CreateTransaction(ctx context.Context, t *libprovider.Transaction) (*libprovider.TransactionResult, error) {
	return nil, libprovider.ErrNotSupported
}
