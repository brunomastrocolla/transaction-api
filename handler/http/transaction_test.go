package http

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/mocks"
)

func TestTransactionHandler(t *testing.T) {

	t.Run("Test Post 1 - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionServiceMock := mocks.NewMockTransactionService(ctrl)
		transactionHandler := NewTransactionHandler(transactionServiceMock)

		req := payloads.TransactionRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          100,
		}

		transactionServiceMock.EXPECT().UpdateBalance(gomock.Any()).Return(nil)
		transactionServiceMock.EXPECT().Create(gomock.Any()).Do(func(transaction *entity.Transaction) {
			transaction.ID = 1
			assert.Equal(t, req.AccountID, transaction.AccountID)
			assert.Equal(t, req.AccountID, transaction.AccountID)
			assert.Equal(t, req.OperationTypeID, transaction.OperationTypeID)
			assert.Equal(t, -req.Amount, transaction.Amount)
			assert.Equal(t, -req.Amount, transaction.Balance)
		}).Return(nil)

		doRequest(t, "POST", "/accounts", req, nil, nil, http.StatusOK, transactionHandler.Post)
	})

	t.Run("Test Post 2 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionServiceMock := mocks.NewMockTransactionService(ctrl)
		transactionHandler := NewTransactionHandler(transactionServiceMock)

		doRequest(t, "POST", "/accounts", "{", nil, nil, http.StatusBadRequest, transactionHandler.Post)
	})

	t.Run("Test Post 3 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionServiceMock := mocks.NewMockTransactionService(ctrl)
		transactionHandler := NewTransactionHandler(transactionServiceMock)

		req := payloads.TransactionRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          100,
		}

		err := errors.New("update-balance")
		transactionServiceMock.EXPECT().UpdateBalance(gomock.Any()).Return(err)

		doRequest(t, "POST", "/accounts", req, nil, nil, http.StatusInternalServerError, transactionHandler.Post)
	})

	t.Run("Test Post 4 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionServiceMock := mocks.NewMockTransactionService(ctrl)
		transactionHandler := NewTransactionHandler(transactionServiceMock)

		req := payloads.TransactionRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          100,
		}

		transactionServiceMock.EXPECT().UpdateBalance(gomock.Any()).Return(nil)
		transactionServiceMock.EXPECT().Create(gomock.Any()).Return(errors.New("create-error"))

		doRequest(t, "POST", "/accounts", req, nil, nil, http.StatusInternalServerError, transactionHandler.Post)
	})
}
