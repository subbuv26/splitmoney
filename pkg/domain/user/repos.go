package user

import "context"

type UserStorer interface {
	CreateUser(context.Context, User) error
	UpdateUser(context.Context, User) error
}
