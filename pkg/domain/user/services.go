package user

import "context"

type UserManager interface {
	SignUp(context.Context, User) error
	SignInWithEmail(_ context.Context, email string) error
	SignInWithPhoneNumber(_ context.Context, phoneNo int) error
	SignOut(_ context.Context, email string) error
}
