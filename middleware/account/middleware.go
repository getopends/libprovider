package libprovider

import "github.com/getopends/libprovider"

type AccountFetcher func(libprovider.AccountFetcher) libprovider.AccountFetcher

type AccountValidator func(libprovider.AccountValidator) libprovider.AccountValidator
