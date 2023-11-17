package examples

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
	}
}

func (p *ExampleProvider) Configure(ctx context.Context, opts *libprovider.Options) error {
	return nil
}
