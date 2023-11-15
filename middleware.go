package libprovider

type (
	TransactionCreatorMiddleware   func(TransactionCreator) TransactionCreator
	TransactionCheckerMiddleware   func(TransactionChecker) TransactionChecker
	TransactionConfirmerMiddleware func(TransactionConfirmer) TransactionConfirmer
	AccountFetcherMiddleware       func(AccountFetcher) AccountFetcher
	AccountValidatorMiddleware     func(AccountValidator) AccountValidator
)
