package libprovider

import (
	"context"
)

// AccountValidator
type AccountValidator interface {
	ValidateAccount(context.Context, *Account) (*ValidateAccountResult, error)
}

// AccountValidatorFunc
type AccountValidatorFunc func(context.Context, *Transaction) (*TransactionResult, error)

// ValidateAccount
func (a AccountValidatorFunc) ValidateAccount(ctx context.Context, transaction *Transaction) (*TransactionResult, error) {
	return a(ctx, transaction)
}

// AccountFetcher
type AccountFetcher interface {
	FetchAccount(context.Context, *Account) (*FetchAccountResult, error)
}

// AccountFetcherFunc
type AccountFetcherFunc func(context.Context, *Transaction) (*TransactionResult, error)

// FetchAccount
func (t AccountFetcherFunc) FetchAccount(ctx context.Context, transaction *Transaction) (*TransactionResult, error) {
	return t(ctx, transaction)
}

// Account
type Account struct{}

// ValidateAccountResult
type ValidateAccountResult struct{}

// FetchAccountResult
type FetchAccountResult struct{}
