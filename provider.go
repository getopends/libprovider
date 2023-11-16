package libprovider

import (
	"context"
	"errors"
)

// ConfigureOptions
type ConfigureOptions struct{}

// Provider
type Provider interface {
	Configuration
	TransactionCreator
	TransactionChecker
	TransactionConfirmer
	AccountFetcher
	AccountValidator
	BalancesFetcher
	ProductsFetcher
}

// NewProviderFunc
type NewProviderFunc func(context.Context) (Provider, error)

// DefaultProvider
type DefaultProvider struct {
	Configuration
	TransactionCreator
	TransactionChecker
	TransactionConfirmer
	AccountFetcher
	AccountValidator
	BalancesFetcher
	ProductsFetcher
}

// Configure
func (d DefaultProvider) Configure(ctx context.Context, opts *ConfigureOptions) error {
	if d.Configuration == nil {
		return errors.ErrUnsupported
	}

	return d.Configuration.Configure(ctx, opts)
}

// CreateTransaction
func (d DefaultProvider) CreateTransaction(ctx context.Context, t *Transaction) (*CreateTransactionResult, error) {
	if d.TransactionCreator == nil {
		return nil, errors.ErrUnsupported
	}

	return d.TransactionCreator.CreateTransaction(ctx, t)
}

// ConfirmTransaction
func (d DefaultProvider) ConfirmTransaction(ctx context.Context, t *Transaction) (*ConfirmTransactionResult, error) {
	if d.TransactionCreator == nil {
		return nil, errors.ErrUnsupported
	}

	return d.TransactionConfirmer.ConfirmTransaction(ctx, t)
}

// CheckTransaction
func (d DefaultProvider) CheckTransaction(ctx context.Context, t *Transaction) (*CheckTransactionResult, error) {
	if d.TransactionCreator == nil {
		return nil, errors.ErrUnsupported
	}

	return d.TransactionChecker.CheckTransaction(ctx, t)
}

// FetchBalances
func (d DefaultProvider) FetchBalances(ctx context.Context, t *Transaction) (*FetchBalancesResult, error) {
	if d.BalancesFetcher == nil {
		return nil, errors.ErrUnsupported
	}

	return d.BalancesFetcher.FetchBalances(ctx, t)
}

// FetchProducts
func (d DefaultProvider) FetchProducts(ctx context.Context, t *Transaction) (*FetchProductsResult, error) {
	if d.ProductsFetcher == nil {
		return nil, errors.ErrUnsupported
	}

	return d.ProductsFetcher.FetchProducts(ctx)
}
