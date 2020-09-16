package mam

import (
	"github.com/muxxer/iota.go/api"
	"github.com/muxxer/iota.go/bundle"
	"github.com/muxxer/iota.go/transaction"
	"github.com/muxxer/iota.go/trinary"
)

// API defines an interface with a subset of methods of `api.API`.
type API interface {
	PrepareTransfers(seed trinary.Trytes, transfers bundle.Transfers, opts api.PrepareTransfersOptions) ([]trinary.Trytes, error)
	SendTrytes(trytes []trinary.Trytes, depth uint64, mwm uint64, reference ...trinary.Hash) (bundle.Bundle, error)
	FindTransactionObjects(query api.FindTransactionsQuery) (transaction.Transactions, error)
}
