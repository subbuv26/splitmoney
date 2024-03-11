package inmem

import (
	"context"
	"splitmoney/pkg/domain/account"
	"splitmoney/pkg/domain/user"
)

type userStore struct {
	phoneIndex     map[int]*user.User
	emailIndex     map[string]*user.User
	accountIDIndex map[account.AccountID]*user.User
}

func NewUserStore() user.UserStorer {
	return &userStore{
		phoneIndex:     make(map[int]*user.User),
		emailIndex:     make(map[string]*user.User),
		accountIDIndex: make(map[account.AccountID]*user.User),
	}
}

func (us *userStore) CreateUser(_ context.Context, u user.User) error {
	us.phoneIndex[u.Phone] = &u
	us.emailIndex[u.Email] = &u
	us.accountIDIndex[u.AccountID] = &u
	return nil
}

func (us *userStore) UpdateUser(_ context.Context, u user.User) error {
	us.phoneIndex[u.Phone] = &u
	us.emailIndex[u.Email] = &u
	us.accountIDIndex[u.AccountID] = &u
	return nil
}

func (us *userStore) GetUser(_ context.Context, accId account.AccountID) (user.User, error) {
	u, ok := us.accountIDIndex[accId]
	if !ok {
		return user.User{}, user.ErrorUserNotFound
	}
	return *u, nil
}

func (us *userStore) GetUserByEmail(_ context.Context, email string) (user.User, error) {
	u, ok := us.emailIndex[email]
	if !ok {
		return user.User{}, user.ErrorUserNotFound
	}
	return *u, nil
}
