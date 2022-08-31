package payloads

import validation "github.com/go-ozzo/ozzo-validation/v4"

type TransactionRequest struct {
	AccountID       int64   `json:"account_id"`
	OperationTypeID int32   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (t TransactionRequest) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.AccountID, validation.Required),
		validation.Field(&t.OperationTypeID, validation.Required),
		validation.Field(&t.Amount, validation.Required),
	)
}
