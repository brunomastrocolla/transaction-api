package repository

import (
	"transaction-api/entity"
)

type AccountRepository interface {
	Create(account *entity.Account) error
	Find(id int32) (entity.Account, error)
}

type MigrationRepository interface {
	Migrate(migrationDir string) error
}

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	Find(id int32) (entity.Transaction, error)
}
