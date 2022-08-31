package postgres

import (
	"github.com/jmoiron/sqlx"

	"transaction-api/entity"
)

type Transaction struct {
	db *sqlx.DB
}

func (t Transaction) Create(transaction *entity.Transaction) error {
	sqlStatement := `
		INSERT INTO transactions (
			"account_id",
			"operation_type_id",
			"amount",
			"created_at",
			"updated_at"
		)
		VALUES (
			:account_id,
			:operation_type_id,
			:amount,
		    :created_at,
			:updated_at
		)
		RETURNING id
	`
	query, err := t.db.PrepareNamed(sqlStatement)
	if err != nil {
		return err
	}
	return query.Get(&transaction.ID, transaction)
}

func (t Transaction) Find(id int32) (entity.Transaction, error) {
	transaction := entity.Transaction{}
	sqlStatement := `
		SELECT *
		FROM transactions
		WHERE id = $1
	`
	err := t.db.Get(&transaction, sqlStatement, id)
	return transaction, err
}

func NewTransactionRepository(db *sqlx.DB) Transaction {
	return Transaction{
		db: db,
	}
}
