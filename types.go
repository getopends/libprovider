package providerd

type (
	Beneficiary struct{}

	Account struct{}

	Balance struct{}

	ValidateAccountResult struct{}

	FetchAccountResult struct{}

	FetchBalancesResult struct{}

	Manifest struct {
		ID      string
		Version string
	}
)
