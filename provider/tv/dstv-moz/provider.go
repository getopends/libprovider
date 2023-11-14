package dstvmoz

import (
	"context"
	"errors"

	"github.com/getopends/providerd"
)

func New(_ context.Context) (providerd.Provider, error) {
	newProvider := &dstvProvider{}

	return &providerd.DefaultProvider{
		TransactionCreator: providerd.TransactionCreatorFunc(newProvider.CreateTransaction),
	}, nil
}

type dstvProvider struct{}

func (d dstvProvider) CreateTransaction(ctx context.Context, t *providerd.Transaction) (*providerd.CreateTransactionResult, error) {
	return nil, errors.ErrUnsupported
}
