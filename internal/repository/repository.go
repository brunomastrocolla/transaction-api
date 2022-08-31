package repository

import (
	entity2 "transaction-api/internal/entity"
)

type AccountRepository interface {
	Create(account *entity2.Account) error
	Find(id int32) (entity2.Account, error)
}

type TransactionRepository interface {
	Create(transaction *entity2.Transaction) error
	Find(id int32) (entity2.Transaction, error)
}
