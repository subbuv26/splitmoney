package auth

import (
	"context"
	"log/slog"
	"splitmoney/pkg/domain/account"
	"splitmoney/pkg/domain/user"
)

type authServer struct {
	repo user.UserStorer
}

func New(userRepo user.UserStorer) user.UserManager {
	return &authServer{
		repo: userRepo,
	}
}

func (as *authServer) SignUp(ctx context.Context, u user.User) (account.AccountID, error) {
	var err error
	if _, err = as.repo.GetUserByEmail(ctx, u.Email); err == nil {
		slog.Error("Email already in use", "user", u)
		return invalidAccountId(), user.ErrorEmailAlreadyInUse
	}
	u.AccountID = getUniqId(u)
	err = as.repo.CreateUser(ctx, u)
	if err != nil {
		return invalidAccountId(), err
	}
	slog.Info("Signed up new User", "user", u)
	return u.AccountID, nil
}

func (as *authServer) SignInWithEmail(_ context.Context, email string) error {
	return nil
}

func (as *authServer) SignInWithPhoneNumber(_ context.Context, phoneNo int) error {
	return nil
}

func (as *authServer) SignOut(_ context.Context, email string) error {
	return nil
}

func getUniqId(u user.User) account.AccountID {
	return account.AccountID(u.Phone)
}

func invalidAccountId() account.AccountID {
	return account.AccountID(-1)
}
