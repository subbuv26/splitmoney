package user

import (
	"context"
	"splitmoney/pkg/domain/account"
)

type UserManager interface {
	SignUp(context.Context, User) (account.AccountID, error)
	SignInWithEmail(_ context.Context, email string) error
	SignInWithPhoneNumber(_ context.Context, phoneNo int) error
	SignOut(_ context.Context, email string) error
}
