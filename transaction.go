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

// TransactionReverter
type TransactionReverter interface {
	RevertTransaction(context.Context, *Transaction) (*RevertTransactionResult, error)
}

// TransactionReverterFunc
type TransactionReverterFunc func(context.Context, *Transaction) (*RevertTransactionResult, error)

// RevertTransaction
func (t TransactionReverterFunc) RevertTransaction(ctx context.Context, transaction *Transaction) (*RevertTransactionResult, error) {
	return t(ctx, transaction)
}

// Transaction
type Transaction struct{}

// CreateTransactionResult
type CreateTransactionResult struct {
	Status Status `json:"status"`
}

// ConfirmTransactionResult
type ConfirmTransactionResult struct{}

// CheckTransactionResult
type CheckTransactionResult struct{}

// RevertTransactionResult
type RevertTransactionResult struct{}

type TransactionStatus uint

var (
	TransactionCreated = Status{
		Code:    1000,
		Message: "CREATED",
		Class:   1,
	}

	TransactionConfirmed = Status{
		Code:    2000,
		Message: "CONFIRMED",
		Class:   2,
	}

	TransactionRejected = Status{
		Code:    3000,
		Message: "REJECTED",
		Class:   3,
	}

	TransactionCancelled = Status{
		Code:    4000,
		Message: "CANCELLED",
		Class:   4,
	}

	TransactionSubmitted = Status{
		Code:    5000,
		Message: "SUBMITTED",
		Class:   5,
	}

	TransactionCompleted = Status{
		Code:    6000,
		Message: "COMPLETED",
		Class:   6,
	}

	TransactionReverted = Status{
		Code:    7000,
		Message: "REVERTED",
		Class:   7,
	}

	TransactionDeclined = Status{
		Code:    8000,
		Message: "DECLINED",
		Class:   8,
	}
)

type Status struct {
	Code    int
	Message string
	Class   int
}
