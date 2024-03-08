package payment

import (
	"context"
	"splitmoney/pkg/domain/account"
)

type Storer interface {
	CreateTx(context.Context, Transaction) error
	GetTxs(context.Context, account.AccountID) ([]Transaction, error)
}
