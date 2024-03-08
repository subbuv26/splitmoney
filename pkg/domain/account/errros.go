package account

import "errors"

var (
	ErrorAccountCreateFailed = errors.New("unable to create account")
	ErrorInvalidAccountData  = errors.New("invalid user data")
	ErrorAccountNotFound     = errors.New("account not found")
)
