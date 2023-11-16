package libprovider

import "github.com/getopends/libprovider"

type TransactionCreator func(libprovider.TransactionCreator) libprovider.TransactionCreator

type TransactionChecker func(libprovider.TransactionChecker) libprovider.TransactionChecker

type TransactionConfirmer func(libprovider.TransactionConfirmer) libprovider.TransactionConfirmer
