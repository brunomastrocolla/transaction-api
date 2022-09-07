package http

import (
	"net/http"

	"go.uber.org/zap"

	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/service"
)

type Transaction struct {
	transactionService service.TransactionService
}

func (t Transaction) Post(w http.ResponseWriter, r *http.Request) {
	transactionRequest := payloads.TransactionRequest{}
	if err := readRequest(r, &transactionRequest); err != nil {
		_ = writeResponse(w, []byte(``), http.StatusBadRequest)
		return
	}

	transaction := entity.Transaction{
		AccountID:       transactionRequest.AccountID,
		OperationTypeID: transactionRequest.OperationTypeID,
	}
	transaction.SetAmount(transactionRequest.Amount)
	transaction.Balance = transaction.Amount

	if err := t.transactionService.UpdateBalance(&transaction); err != nil {
		zap.L().Error("update-balance-error", zap.Error(err))
		_ = writeResponse(w, []byte(``), http.StatusInternalServerError)
		return
	}

	if err := t.transactionService.Create(&transaction); err != nil {
		zap.L().Error("transaction-create-error", zap.Error(err))
		_ = writeResponse(w, []byte(``), http.StatusInternalServerError)
		return
	}

	zap.L().Info("transaction-created", zap.Any("transaction", transaction))
	_ = writeResponse(w, []byte(``), http.StatusOK)
}

func NewTransactionHandler(transactionService service.TransactionService) Transaction {
	return Transaction{
		transactionService: transactionService,
	}
}
