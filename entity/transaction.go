package entity

import "time"

type Transaction struct {
	ID              int64     `db:"id"`
	AccountID       int64     `db:"account_id"`
	OperationTypeID int32     `db:"operation_type_id"`
	Amount          float64   `db:"amount"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func (t *Transaction) SetAmount(amount float64) {
	switch t.OperationTypeID {
	case 1, 2, 3:
		t.Amount = -amount
	case 4:
		t.Amount = amount
	}
}
