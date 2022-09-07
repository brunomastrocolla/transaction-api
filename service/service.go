package service

import "transaction-api/entity"

type AccountService interface {
	Create(account *entity.Account) error
	Find(id int64) (entity.Account, error)
}

type TransactionService interface {
	Create(transaction *entity.Transaction) error
	Find(id int32) (entity.Transaction, error)
	FindByAccountID(id int64) ([]entity.Transaction, error)
	Update(transaction *entity.Transaction) error
	UpdateBalance(transaction *entity.Transaction) error
}
