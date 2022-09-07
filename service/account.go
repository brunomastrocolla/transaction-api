package service

import (
	"time"

	"transaction-api/entity"
	"transaction-api/repository"
)

type Account struct {
	accountRepo repository.AccountRepository
}

func (a Account) Create(account *entity.Account) error {
	now := time.Now()
	account.CreatedAt = now
	account.UpdatedAt = now
	return a.accountRepo.Create(account)
}

func (a Account) Find(id int64) (entity.Account, error) {
	return a.accountRepo.Find(id)
}

func NewAccountService(accountRepo repository.AccountRepository) Account {
	return Account{
		accountRepo: accountRepo,
	}
}
