package payment

import "splitmoney/pkg/domain/account"

type Transaction struct {
	FromAccount account.AccountID
	ToAccount   account.AccountID
	Amount      int
}
