package myapp

// NewAccount constructs new account
func NewAccount(holder string) (*Account, error) {
	return &Account{
		ID: "123456",
	}, nil
}

// Account represents bank account domain aggregate
type Account struct {
	ID string
}
