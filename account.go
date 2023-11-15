package libprovider

import (
	"context"
)

type (
	Account struct{}

	ValidateAccountResult struct{}

	FetchAccountResult struct{}

	FetchBalancesResult struct{}
)

type AccountValidator interface {
	ValidateAccount(context.Context, *Account) (*ValidateAccountResult, error)
}

type AccountValidatorFunc func(context.Context, *Transaction) (*ConfirmTransactionResult, error)

func (a AccountValidatorFunc) ValidateAccount(ctx context.Context, transaction *Transaction) (*ConfirmTransactionResult, error) {
	return a(ctx, transaction)
}

type AccountFetcher interface {
	FetchAccount(context.Context, *Account) (*FetchAccountResult, error)
}

type AccountFetcherFunc func(context.Context, *Transaction) (*ConfirmTransactionResult, error)

func (t AccountFetcherFunc) FetchAccount(ctx context.Context, transaction *Transaction) (*ConfirmTransactionResult, error) {
	return t(ctx, transaction)
}
