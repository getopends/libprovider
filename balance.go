package providerd

import (
	"context"
)

type BalanceFetcher interface {
	FetchBalances(context.Context, *Transaction) (*FetchBalancesResult, error)
}

type BalanceFetcherFunc func(context.Context, *Transaction) (*FetchBalancesResult, error)

func (b BalanceFetcherFunc) FetchBalances(ctx context.Context, transaction *Transaction) (*FetchBalancesResult, error) {
	return b(ctx, transaction)
}
