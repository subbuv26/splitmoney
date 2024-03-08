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
