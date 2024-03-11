package inmem

import (
	"context"
	"splitmoney/pkg/domain/account"
)

type accountStore struct {
	accounts map[account.AccountID]*account.Account
}

func NewAccountStore() account.Storer {
	return &accountStore{
		accounts: make(map[account.AccountID]*account.Account),
	}
}

func (as *accountStore) CreateAccount(_ context.Context, acc account.Account) error {
	as.accounts[acc.ID] = &acc
	return nil
}

func (as *accountStore) GetAccount(_ context.Context, accId account.AccountID) (account.Account, error) {
	acc, ok := as.accounts[accId]
	if !ok {
		return account.Account{}, account.ErrorAccountNotFound
	}
	return *acc, nil
}
