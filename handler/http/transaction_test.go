package http

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/mocks"
)

func doRequest(t *testing.T, method, url string, request interface{}, response interface{},
	handler func(w http.ResponseWriter, r *http.Request)) {

	reqBody, err := json.Marshal(request)
	assert.NilError(t, err)

	httpReq, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	assert.NilError(t, err)

	httpRec := httptest.NewRecorder()
	handler(httpRec, httpReq)

	reqResult := httpRec.Result()
	assert.Equal(t, http.StatusOK, reqResult.StatusCode)

	resBody, err := ioutil.ReadAll(reqResult.Body)
	assert.NilError(t, err)

	err = json.Unmarshal(resBody, &response)
	assert.NilError(t, err)
}

func TestTransaction(t *testing.T) {

	t.Run("Test Transaction 1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionHandler := NewTransactionHandler(transactionRepoMock)

		transactionRepoMock.EXPECT().FindByAccountID(gomock.Any()).Return(
			[]entity.Transaction{
				{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -50,
					Balance:         -50,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              2,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -23.5,
					Balance:         -23.5,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              3,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -18.7,
					Balance:         -18.7,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
			}, nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(-13.5), transaction.Balance)
		}).Return(nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(-18.7), transaction.Balance)
		}).Return(nil)
		transactionRepoMock.EXPECT().Create(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)

		trans := payloads.TransactionRequest{
			AccountID:       1,
			OperationTypeID: 4,
			Amount:          60,
		}

		doRequest(t, "POST", "", trans, nil, transactionHandler.Post)

	})

	t.Run("Test Transaction 2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionHandler := NewTransactionHandler(transactionRepoMock)

		transactionRepoMock.EXPECT().FindByAccountID(gomock.Any()).Return(
			[]entity.Transaction{
				{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -50,
					Balance:         0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              2,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -23.5,
					Balance:         -13.5,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              3,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -18.7,
					Balance:         -18.7,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              4,
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          60,
					Balance:         0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
			}, nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)
		/*transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)*/
		transactionRepoMock.EXPECT().Create(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(67.8), transaction.Balance)
		}).Return(nil)

		trans := payloads.TransactionRequest{
			AccountID:       1,
			OperationTypeID: 4,
			Amount:          100,
		}

		doRequest(t, "POST", "", trans, nil, transactionHandler.Post)

	})
}
