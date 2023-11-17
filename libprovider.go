package libprovider

import (
	"context"
	"errors"
)

var (
	// ErrProviderNotFound
	ErrProviderNotFound = errors.New("not found")

	// ErrNotSupported
	ErrNotSupported = errors.New("not supported")

	// ErrProviderDuplicated
	ErrProviderDuplicated = errors.New("provider duplicated")
)

type Options struct{}

type Info struct {
	Code    string
	Name    string
	Slug    string
	Version string
	Author  string
}

// Provider
type Provider interface {
	Info() Info
	Configure(context.Context, *Options) error
}

// NewProviderFunc
type NewProviderFunc func(context.Context) Provider

// DefaultProvider
type DefaultProvider struct {
	TransactionCreator
	TransactionChecker
	TransactionConfirmer
	AccountFetcher
	AccountValidator
	BalancesFetcher
	ProductsFetcher
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
