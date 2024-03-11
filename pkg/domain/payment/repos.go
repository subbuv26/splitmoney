package payment

import (
	"context"
	"splitmoney/pkg/domain/account"
	"time"
)

type Storer interface {
	CreateTx(context.Context, Transaction) error
	GetTxs(context.Context, account.AccountID) ([]Transaction, error)
	GetTxsInDuration(_ context.Context, accID account.AccountID, from time.Time, to time.Time) ([]Transaction, error)
}
