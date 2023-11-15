package providerd

import (
	"context"
)

type (
	Transaction struct {
		ID *int64 `json:"id"`
	}

	CreateTransactionResult struct{}

	ConfirmTransactionResult struct{}

	CheckTransactionResult struct{}
)

type TransactionCreator interface {
	CreateTransaction(context.Context, *Transaction) (*CreateTransactionResult, error)
}

type TransactionCreatorFunc func(context.Context, *Transaction) (*CreateTransactionResult, error)

func (t TransactionCreatorFunc) CreateTransaction(ctx context.Context, transaction *Transaction) (*CreateTransactionResult, error) {
	return t(ctx, transaction)
}

type TransactionConfirmer interface {
	ConfirmTransaction(context.Context, *Transaction) (*ConfirmTransactionResult, error)
}

type TransactionConfirmerFunc func(context.Context, *Transaction) (*ConfirmTransactionResult, error)

func (t TransactionConfirmerFunc) ConfirmTransaction(ctx context.Context, transaction *Transaction) (*ConfirmTransactionResult, error) {
	return t(ctx, transaction)
}

type TransactionChecker interface {
	CheckTransaction(context.Context, *Transaction) (*CheckTransactionResult, error)
}

type TransactionCheckerFunc func(context.Context, *Transaction) (*CheckTransactionResult, error)

func (t TransactionCheckerFunc) CheckTransaction(ctx context.Context, transaction *Transaction) (*CheckTransactionResult, error) {
	return t(ctx, transaction)
}
