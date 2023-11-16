package libprovider

import (
	"context"
)

// BalancesFetcher
type BalancesFetcher interface {
	FetchBalances(context.Context, *Transaction) (*FetchBalancesResult, error)
}

// BalancesFetcherFunc
type BalancesFetcherFunc func(context.Context, *Transaction) (*FetchBalancesResult, error)

// FetchBalances
func (b BalancesFetcherFunc) FetchBalances(ctx context.Context, transaction *Transaction) (*FetchBalancesResult, error) {
	return b(ctx, transaction)
}

// Balance
type Balance struct{}

// FetchBalancesResult
type FetchBalancesResult struct{}
