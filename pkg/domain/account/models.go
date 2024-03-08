package account

type AccountID int

type Account struct {
	ID      AccountID
	Debt    int
	Lend    int
	Balance int
}
