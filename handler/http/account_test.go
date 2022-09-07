package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/mocks"
)

func addURLParam(param map[string]string, httpReq *http.Request) *http.Request {
	if len(param) == 0 {
		return httpReq
	}
	chiContext := chi.NewRouteContext()
	for k, v := range param {
		chiContext.URLParams.Add(k, v)
	}
	return httpReq.WithContext(context.WithValue(httpReq.Context(), chi.RouteCtxKey, chiContext))
}

func doRequest(t *testing.T, method, url string, request interface{}, response interface{},
	param map[string]string, statusCode int, handler func(w http.ResponseWriter, r *http.Request)) {

	reqBody, err := json.Marshal(request)
	assert.NilError(t, err)

	httpReq, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	assert.NilError(t, err)

	httpReq = addURLParam(param, httpReq)

	httpRec := httptest.NewRecorder()
	handler(httpRec, httpReq)

	reqResult := httpRec.Result()
	assert.Equal(t, statusCode, reqResult.StatusCode)

	resBody, err := ioutil.ReadAll(reqResult.Body)
	assert.NilError(t, err)

	err = json.Unmarshal(resBody, &response)
	assert.NilError(t, err)
}

func TestAccountHandler(t *testing.T) {

	t.Run("Test Post 1 - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountServiceMock := mocks.NewMockAccountService(ctrl)
		accountHandler := NewAccountHandler(accountServiceMock)

		req := payloads.AccountRequest{DocumentNumber: "01234567890"}
		res := payloads.AccountResponse{}

		accountServiceMock.EXPECT().Create(gomock.Any()).Do(func(account *entity.Account) {
			account.ID = 1
			assert.Equal(t, req.DocumentNumber, account.DocumentNumber)
		}).Return(nil)

		doRequest(t, "POST", "/accounts", req, &res, nil, http.StatusOK, accountHandler.Post)
		assert.Equal(t, int64(1), res.AccountID)
		assert.Equal(t, req.DocumentNumber, res.DocumentNumber)
	})

	t.Run("Test Post 2 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountServiceMock := mocks.NewMockAccountService(ctrl)
		accountHandler := NewAccountHandler(accountServiceMock)

		doRequest(t, "POST", "/accounts", "{", nil, nil, http.StatusBadRequest, accountHandler.Post)
	})

	t.Run("Test Post 3 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountServiceMock := mocks.NewMockAccountService(ctrl)
		accountHandler := NewAccountHandler(accountServiceMock)

		req := payloads.AccountRequest{DocumentNumber: "01234567890"}

		err := errors.New("create-error")
		accountServiceMock.EXPECT().Create(gomock.Any()).Return(err)

		doRequest(t, "POST", "/accounts", req, nil, nil, http.StatusInternalServerError, accountHandler.Post)
	})

	t.Run("Test Get 1 - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountServiceMock := mocks.NewMockAccountService(ctrl)
		accountHandler := NewAccountHandler(accountServiceMock)

		accountServiceMock.EXPECT().Find(int64(1)).Return(entity.Account{
			ID:             1,
			DocumentNumber: "01234567890",
		}, nil)

		res := payloads.AccountResponse{}
		params := map[string]string{
			"id": "1",
		}

		doRequest(t, "POST", "/accounts", nil, &res, params, http.StatusOK, accountHandler.Get)

		assert.Equal(t, int64(1), res.AccountID)
		assert.Equal(t, "01234567890", res.DocumentNumber)
	})

	t.Run("Test Get 2 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountServiceMock := mocks.NewMockAccountService(ctrl)
		accountHandler := NewAccountHandler(accountServiceMock)

		params := map[string]string{
			"id": "abc",
		}

		doRequest(t, "POST", "/accounts", nil, nil, params, http.StatusBadRequest, accountHandler.Get)
	})

	t.Run("Test Get 3 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountServiceMock := mocks.NewMockAccountService(ctrl)
		accountHandler := NewAccountHandler(accountServiceMock)

		err := errors.New("find-error")
		accountServiceMock.EXPECT().Find(int64(1)).Return(entity.Account{}, err)

		params := map[string]string{
			"id": "1",
		}

		doRequest(t, "POST", "/accounts", nil, nil, params, http.StatusNotFound, accountHandler.Get)
	})

}
