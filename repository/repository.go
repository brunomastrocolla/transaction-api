package repository

import (
	"transaction-api/entity"
)

type MigrationType string

var (
	MigrateUp   MigrationType = "up"
	MigrateDown MigrationType = "down"
	MigrateDrop MigrationType = "drop"
)

type AccountRepository interface {
	Create(account *entity.Account) error
	Find(id int64) (entity.Account, error)
}

type MigrationRepository interface {
	Migrate(migrationDir string, migrationType MigrationType) error
}

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	Find(id int64) (entity.Transaction, error)
	FindByAccountID(id int64) ([]entity.Transaction, error)
	Update(transaction *entity.Transaction) error
}
