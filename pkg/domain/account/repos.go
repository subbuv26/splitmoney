package account

import "context"

type Storer interface {
	CreateAccount(context.Context, Account) error
	GetAccount(context.Context, AccountID) (Account, error)
}
