package providerd

import (
	"context"
	"errors"
)

type DefaultProvider struct {
	Configurer
	TransactionCreator
	TransactionChecker
	TransactionConfirmer
	AccountFetcher
	AccountValidator
	BalanceFetcher
}

func (d DefaultProvider) Configure(ctx context.Context) error {
	if d.Configurer == nil {
		return errors.ErrUnsupported
	}

	return d.Configurer.Configure(ctx)
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
	if d.BalanceFetcher == nil {
		return nil, errors.ErrUnsupported
	}

	return d.BalanceFetcher.FetchBalances(ctx, t)
}

type Provider interface {
	Configurer
	TransactionCreator
	TransactionChecker
	TransactionConfirmer
	AccountFetcher
	AccountValidator
	BalanceFetcher
}

type NewProviderFunc func(context.Context) (Provider, error)

func CheckCountry(tc TransactionCreator, countries ...string) TransactionCreator {
	return TransactionCreatorFunc(func(ctx context.Context, t *Transaction) (*CreateTransactionResult, error) {
		return tc.CreateTransaction(ctx, t)
	})
}
