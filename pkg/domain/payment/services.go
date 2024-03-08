package payment

import "context"

type PaymentsManager interface {
	MakeTx(context.Context, Transaction) error
}
