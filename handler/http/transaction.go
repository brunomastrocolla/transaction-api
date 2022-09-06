package http

import (
	"go.uber.org/zap"
	"net/http"
	"time"
	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/repository"
)

type Transaction struct {
	transactionRepo repository.TransactionRepository
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
	transaction.Balance = transaction.Amount

	if err := t.updateBalance(&transaction); err != nil {
		zap.L().Error("update-balance-error", zap.Error(err))
		writeResponse(w, []byte(``), http.StatusInternalServerError)
		return
	}

	if err := t.transactionRepo.Create(&transaction); err != nil {
		zap.L().Error("transaction-create-error", zap.Error(err))
		writeResponse(w, []byte(``), http.StatusInternalServerError)
		return
	}

	zap.L().Info("transaction-created", zap.Any("transaction", transaction))

	writeResponse(w, []byte(``), http.StatusOK)
}

func (t Transaction) updateBalance(transaction *entity.Transaction) error {
	if transaction.OperationTypeID == 4 {
		trans, err := t.transactionRepo.FindByAccountID(transaction.AccountID)
		if err != nil {
			return err
		}

		for i, _ := range trans {
			a := &trans[i]

			if transaction.Balance == 0 {
				break
			}

			switch a.OperationTypeID {
			case 1, 2, 3:
				balance := -a.Balance
				if balance == 0 {
					continue
				}
				if balance > transaction.Balance {
					a.Balance = -(balance - transaction.Balance)
					transaction.Balance = 0
				} else {
					a.Balance = 0
					transaction.Balance -= balance
				}
			}
		}

		for _, tt := range trans {
			if tt.OperationTypeID != 4 {
				err := t.transactionRepo.Update(&tt)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewTransactionHandler(transactionRepo repository.TransactionRepository) Transaction {
	return Transaction{
		transactionRepo: transactionRepo,
	}
}
