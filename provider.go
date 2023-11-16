package libprovider

import (
	"context"
	"errors"
)

type Options struct{}

type Provider interface {
	Configurer
	TransactionCreator
	TransactionChecker
	TransactionConfirmer
	AccountFetcher
	AccountValidator
	BalancesFetcher
}

type NewProviderFunc func(context.Context) (Provider, error)

type DefaultProvider struct {
	Configurer
	TransactionCreator
	TransactionChecker
	TransactionConfirmer
	AccountFetcher
	AccountValidator
	BalancesFetcher
}

func (d DefaultProvider) Configure(ctx context.Context, opts *Options) error {
	if d.Configurer == nil {
		return errors.ErrUnsupported
	}

	return d.Configurer.Configure(ctx, opts)
}

func (d DefaultProvider) CreateTransaction(ctx context.Context, t *Transaction) (*CreateTransactionResult, error) {
	if d.TransactionCreator == nil {
		return nil, errors.ErrUnsupported
	}

	return d.TransactionCreator.CreateTransaction(ctx, t)
}

func (d DefaultProvider) ConfirmTransaction(ctx context.Context, t *Transaction) (*ConfirmTransactionResult, error) {
	if d.TransactionCreator == nil {
		return nil, errors.ErrUnsupported
	}

	return d.TransactionConfirmer.ConfirmTransaction(ctx, t)
}

func (d DefaultProvider) CheckTransaction(ctx context.Context, t *Transaction) (*CheckTransactionResult, error) {
	if d.TransactionCreator == nil {
		return nil, errors.ErrUnsupported
	}

	return d.TransactionChecker.CheckTransaction(ctx, t)
}

func (d DefaultProvider) FetchBalances(ctx context.Context, t *Transaction) (*FetchBalancesResult, error) {
	if d.BalancesFetcher == nil {
		return nil, errors.ErrUnsupported
	}

	return d.BalancesFetcher.FetchBalances(ctx, t)
}
