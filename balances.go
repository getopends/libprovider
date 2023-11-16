package libprovider

import (
	"context"
)

type BalancesFetcher interface {
	FetchBalances(context.Context, *Transaction) (*FetchBalancesResult, error)
}

type BalancesFetcherFunc func(context.Context, *Transaction) (*FetchBalancesResult, error)

func (b BalancesFetcherFunc) FetchBalances(ctx context.Context, transaction *Transaction) (*FetchBalancesResult, error) {
	return b(ctx, transaction)
}

type Balance struct{}

type FetchBalancesResult struct{}
