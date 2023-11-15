package providerd

type (
	Transaction struct{}

	Beneficiary struct{}

	Account struct{}

	Balance struct{}

	CreateTransactionResult struct{}

	ConfirmTransactionResult struct{}

	CheckTransactionResult struct{}

	ValidateAccountResult struct{}

	FetchAccountResult struct{}

	FetchBalancesResult struct{}

	Manifest struct {
		ID      string
		Version string
	}
)
