package user

import (
	"splitmoney/pkg/domain/account"
)

type User struct {
	AccountID account.AccountID
	Profile   Profile
	Phone     int
	Email     string
}

type Profile struct {
	Name    string
	Address string
}
