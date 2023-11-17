package example

import (
	"context"

	"github.com/getopends/libprovider"
)

var _ libprovider.Provider = &ExampleProvider{}

const (
	envHost     = "HOST"
	envPort     = "PORT"
	envUsername = "USERNAME"
	envpassword = "PASSWORD"
)

func NewExampleProvider(context.Context) libprovider.Provider {
	return &ExampleProvider{}
}

type ExampleProvider struct {
	opts             *libprovider.Options
	apiHost, apiPort string
}

func (p *ExampleProvider) Info() libprovider.Info {
	return libprovider.Info{
		Name:    "Example",
		Slug:    "example-moz",
		Author:  "Edson Michaque",
		Version: "0.1.0",
		Variables: map[string]libprovider.VariableDefinition{
			"HOST": {
				Default: libprovider.String("example.com"),
			},
			"PORT": {
				Default: libprovider.String("10"),
			},
			"USERNAME": {
				Secret: true,
			},
			"PASSWORD": {
				Secret: true,
			},
		},
	}
}

func (p *ExampleProvider) Configure(ctx context.Context, opts *libprovider.Options) error {
	apiHost, err := p.opts.Variables.Get(envHost)
	if err != nil {
		return err
	}

	apiPort, err := p.opts.Variables.Get(envPort)
	if err != nil {
		return err
	}

	_, err = p.opts.Variables.Get(envUsername)
	if err != nil {
		return err
	}

	_, err = p.opts.Variables.Get(envPassword)
	if err != nil {
		return err
	}

	p.apiHost = apiHost
	p.apiPort = apiPort

	return nil
}

func (p *ExampleProvider) CreateTransaction(ctx context.Context, t *libprovider.Transaction) (*libprovider.TransactionResult, error) {
	return nil, libprovider.ErrNotSupported
}
