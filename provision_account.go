package myapp

import "context"

// ProvisionAccountFunc represents account provisioning use case
type ProvisionAccountFunc func(context.Context, AccountToProvision) (*ProvisionedAccount, error)

// AccountToProvision holds details of an account to provision
type AccountToProvision struct {
	AccountHolder string
	HolderAddress string
}

// ProvisionedAccount holds provisioned account information
type ProvisionedAccount struct {
	AccountID string
}

// AccountStore represents account store
type AccountStore interface {
	Save(context.Context, Account) error
}

// NewProvisionAccount instantiates new account provisioning use case
func NewProvisionAccount(store AccountStore) ProvisionAccountFunc {
	return func(ctx context.Context, accToProvision AccountToProvision) (*ProvisionedAccount, error) {
		account, err := NewAccount(accToProvision.AccountHolder)
		if err != nil {
			return nil, err
		}

		err = store.Save(ctx, *account)
		if err != nil {
			return nil, err
		}

		return &ProvisionedAccount{
			AccountID: account.ID,
		}, nil
	}
}
