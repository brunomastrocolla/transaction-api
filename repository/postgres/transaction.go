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
		    "balance",
			"created_at",
			"updated_at"
		)
		VALUES (
			:account_id,
			:operation_type_id,
			:amount,
		    :balance,
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

func (t Transaction) FindByAccountID(id int64) ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	sqlStatement := `
		SELECT *
		FROM transactions
		WHERE account_id = $1
		ORDER BY created_at ASC
	`
	err := t.db.Select(&transaction, sqlStatement, id)
	return transaction, err
}

func (t Transaction) Update(transaction *entity.Transaction) error {
	sqlStatement := `
		UPDATE transactions
		SET	account_id = :account_id,
			operation_type_id = :operation_type_id,
			amount = :amount,
			balance = :balance,
			created_at = :created_at,
			updated_at = :updated_at
		WHERE id = :id
	`
	_, err := t.db.NamedExec(sqlStatement, transaction)
	return err
}

func NewTransactionRepository(db *sqlx.DB) Transaction {
	return Transaction{
		db: db,
	}
}
