package http

import (
	"net/http"
	"time"
	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/repository/postgres"

	"go.uber.org/zap"
)

type Transaction struct {
	transactionRepo postgres.Transaction
}

func (t Transaction) Post(w http.ResponseWriter, r *http.Request) {
	transactionRequest := payloads.TransactionRequest{}
	if err := readRequest(r, &transactionRequest); err != nil {
		writeResponse(w, []byte(``), http.StatusBadRequest)
		return
	}

	now := time.Now()
	transaction := entity.Transaction{
		AccountID:       transactionRequest.AccountID,
		OperationTypeID: transactionRequest.OperationTypeID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	transaction.SetAmount(transactionRequest.Amount)

	if err := t.transactionRepo.Create(&transaction); err != nil {
		zap.L().Error("transaction-create-error", zap.Error(err))
		writeResponse(w, []byte(``), http.StatusInternalServerError)
		return
	}
	zap.L().Info("transaction-created", zap.Any("transaction", transaction))

	writeResponse(w, []byte(``), http.StatusOK)
}

func NewTransactionHandler(transactionRepo postgres.Transaction) Transaction {
	return Transaction{
		transactionRepo: transactionRepo,
	}
}
