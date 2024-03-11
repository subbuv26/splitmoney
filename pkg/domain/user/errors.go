package user

import "errors"

var (
	ErrorUserNotFound      = errors.New("user not found")
	ErrorEmailAlreadyInUse = errors.New("email alraedy in use")
	ErrorFailedToSignup    = errors.New("failed to signup")
)
