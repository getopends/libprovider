package libprovider

import (
	"context"
)

type ProductsFetcher interface {
	FetchProducts(context.Context) (*FetchProductsResult, error)
}

type ProductsFetcherFunc func(context.Context) (*FetchProductsResult, error)

func (b ProductsFetcherFunc) FetchProducts(ctx context.Context) (*FetchProductsResult, error) {
	return b(ctx)
}

type Product struct{}

type FetchProductsResult struct {
	Products []Product
}
