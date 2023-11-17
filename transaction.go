package libprovider

import (
	"context"
)

// TransactionCreator
type TransactionCreator interface {
	CreateTransaction(context.Context, *Transaction) (*CreateTransactionResult, error)
}

// TransactionCreatorFunc
type TransactionCreatorFunc func(context.Context, *Transaction) (*CreateTransactionResult, error)

// CreateTransaction
func (t TransactionCreatorFunc) CreateTransaction(ctx context.Context, transaction *Transaction) (*CreateTransactionResult, error) {
	return t(ctx, transaction)
}

// TransactionConfirmer
type TransactionConfirmer interface {
	ConfirmTransaction(context.Context, *Transaction) (*ConfirmTransactionResult, error)
}

// TransactionConfirmerFunc
type TransactionConfirmerFunc func(context.Context, *Transaction) (*ConfirmTransactionResult, error)

// ConfirmTransaction
func (t TransactionConfirmerFunc) ConfirmTransaction(ctx context.Context, transaction *Transaction) (*ConfirmTransactionResult, error) {
	return t(ctx, transaction)
}

// TransactionChecker
type TransactionChecker interface {
	CheckTransaction(context.Context, *Transaction) (*CheckTransactionResult, error)
}

// TransactionCheckerFunc
type TransactionCheckerFunc func(context.Context, *Transaction) (*CheckTransactionResult, error)

// CheckTransaction
func (t TransactionCheckerFunc) CheckTransaction(ctx context.Context, transaction *Transaction) (*CheckTransactionResult, error) {
	return t(ctx, transaction)
}

// Transaction
type Transaction struct{}

// CreateTransactionResult
type CreateTransactionResult struct {
	Status TransactionStatus `json:"status"`
}

// ConfirmTransactionResult
type ConfirmTransactionResult struct{}

// CheckTransactionResult
type CheckTransactionResult struct{}

type TransactionStatus uint8

const (
	TransactionAccepted TransactionStatus = iota
	TransactionDeclided
	TransactionCompleted
)
