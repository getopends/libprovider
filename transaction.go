package libprovider

import (
	"context"
)

type TransactionStatus uint

const (
	TransactionCreated   = TransactionStatus(1000)
	TransactionConfirmed = TransactionStatus(2000)
	TransactionRejected  = TransactionStatus(3000)
	TransactionCancelled = TransactionStatus(4000)
	TransactionSubmitted = TransactionStatus(5000)
	TransactionCompleted = TransactionStatus(6000)
	TransactionReverted  = TransactionStatus(7000)
	TransactionDeclined  = TransactionStatus(8000)
)

// TransactionCreator
type TransactionCreator interface {
	CreateTransaction(context.Context, *Transaction) (*TransactionResult, error)
}

// TransactionCreatorFunc
type TransactionCreatorFunc func(context.Context, *Transaction) (*TransactionResult, error)

// CreateTransaction
func (t TransactionCreatorFunc) CreateTransaction(ctx context.Context, transaction *Transaction) (*TransactionResult, error) {
	return t(ctx, transaction)
}

// TransactionConfirmer
type TransactionConfirmer interface {
	ConfirmTransaction(context.Context, *Transaction) (*TransactionResult, error)
}

// TransactionConfirmerFunc
type TransactionConfirmerFunc func(context.Context, *Transaction) (*TransactionResult, error)

// ConfirmTransaction
func (t TransactionConfirmerFunc) ConfirmTransaction(ctx context.Context, transaction *Transaction) (*TransactionResult, error) {
	return t(ctx, transaction)
}

// TransactionChecker
type TransactionChecker interface {
	CheckTransaction(context.Context, *Transaction) (*TransactionResult, error)
}

// TransactionCheckerFunc
type TransactionCheckerFunc func(context.Context, *Transaction) (*TransactionResult, error)

// CheckTransaction
func (t TransactionCheckerFunc) CheckTransaction(ctx context.Context, transaction *Transaction) (*TransactionResult, error) {
	return t(ctx, transaction)
}

// TransactionReverter
type TransactionReverter interface {
	RevertTransaction(context.Context, *Transaction) (*TransactionResult, error)
}

// TransactionReverterFunc
type TransactionReverterFunc func(context.Context, *Transaction) (*TransactionResult, error)

// RevertTransaction
func (t TransactionReverterFunc) RevertTransaction(ctx context.Context, transaction *Transaction) (*TransactionResult, error) {
	return t(ctx, transaction)
}

// Transaction
type Transaction struct {
	ProductID       *int64          `json:"product_id"`
	ReceivingMethod ReceivingMethod `json:"receiving_method"`
	Operator        Operator        `json:"operator"`
}

type Operator struct {
	ID int64 `json:"id"`
}

type ReceivingMethod struct {
	CardNumber    string `json:"card_number"`
	PhoneNumber   string `json:"phone_number"`
	AccountNumber string `json:"account_number"`
}

// TransactionResult
type TransactionResult struct {
	Status TransactionStatus `json:"status"`
}
