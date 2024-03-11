package user

import (
	"context"
	"splitmoney/pkg/domain/account"
)

type UserStorer interface {
	CreateUser(context.Context, User) error
	UpdateUser(context.Context, User) error
	GetUser(context.Context, account.AccountID) (User, error)
	GetUserByEmail(context.Context, string) (User, error)
}
