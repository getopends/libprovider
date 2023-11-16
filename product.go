package libprovider

import (
	"context"
)

type ProductFetcher interface {
	FetchProducts(context.Context) (*FetchProductsResult, error)
}

type ProductFetcherFunc func(context.Context) (*FetchProductsResult, error)

func (b ProductFetcherFunc) FetchProducts(ctx context.Context) (*FetchProductsResult, error) {
	return b(ctx)
}

type Product struct{}

type FetchProductsResult struct {
	Products []Product
}
