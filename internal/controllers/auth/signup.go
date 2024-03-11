package auth

import "splitmoney/pkg/domain/account"

type SignupRequest struct {
	Email string `json:"email"`
	Phone int    `json:"phone"`
	Name  string `json:"name"`
}

type SingupResponse struct {
	AccountID account.AccountID `json:"accountID"`
}
