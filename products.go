package libprovider

import (
	"context"
)

// ProductsFetcher
type ProductsFetcher interface {
	FetchProducts(context.Context) (*FetchProductsResult, error)
}

// ProductsFetcherFunc
type ProductsFetcherFunc func(context.Context) (*FetchProductsResult, error)

// FetchProducts
func (p ProductsFetcherFunc) FetchProducts(ctx context.Context) (*FetchProductsResult, error) {
	return p(ctx)
}

// Product
type Product struct{}

// FetchProductsResult
type FetchProductsResult struct {
	Products []Product
}
