package redis

import (
	"splitmoney/pkg/domain/account"
	"time"
)

type accountModel struct {
	ID        account.AccountID `json:"id"`
	Debt      int               `json:"debt"`
	Lend      int               `json:"lend"`
	Balance   int               `json:"balance"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

type paymentModel struct {
	FromAccount   account.AccountID `json:"fromAccount"`
	ToAccount     account.AccountID `json:"toAccount"`
	Ammount       int               `json:"ammount"`
	TransactionId int               `json:"TrasnsactionId"`
	CreatedAt     time.Time
}
