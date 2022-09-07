package postgres

import (
	"github.com/jmoiron/sqlx"

	"transaction-api/entity"
)

type Account struct {
	db *sqlx.DB
}

func (a Account) Create(account *entity.Account) error {
	sqlStatement := `
		INSERT INTO accounts (
			"document_number",
			"created_at",
			"updated_at"
		)
		VALUES (
			:document_number,
			:created_at,
			:updated_at
		)
		RETURNING id
	`
	query, err := a.db.PrepareNamed(sqlStatement)
	if err != nil {
		return err
	}
	return query.Get(&account.ID, account)
}

func (a Account) Find(id int64) (entity.Account, error) {
	account := entity.Account{}
	sqlStatement := `
		SELECT *
		FROM accounts
		WHERE id = $1
	`
	err := a.db.Get(&account, sqlStatement, id)
	return account, err
}

func NewAccountRepository(db *sqlx.DB) Account {
	return Account{
		db: db,
	}
}
