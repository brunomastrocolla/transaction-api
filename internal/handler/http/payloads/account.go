package payloads

import validation "github.com/go-ozzo/ozzo-validation/v4"

type AccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (a AccountRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.DocumentNumber, validation.Required),
	)
}

type AccountResponse struct {
	AccountID      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
