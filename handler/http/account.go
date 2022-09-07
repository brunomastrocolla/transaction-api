package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/service"
)

type Account struct {
	accountService service.AccountService
}

func (a Account) Post(w http.ResponseWriter, r *http.Request) {
	accountRequest := payloads.AccountRequest{}
	if err := readRequest(r, &accountRequest); err != nil {
		_ = writeResponse(w, []byte(``), http.StatusBadRequest)
		return
	}

	account := entity.Account{
		DocumentNumber: accountRequest.DocumentNumber,
	}

	if err := a.accountService.Create(&account); err != nil {
		zap.L().Error("account-store-error", zap.Error(err))
		_ = writeResponse(w, []byte(``), http.StatusInternalServerError)
		return
	}
	zap.L().Info("account-created", zap.Any("account", account))

	accountResponse := payloads.AccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
	_ = writeResponse(w, accountResponse, http.StatusOK)
}

func (a Account) Get(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		zap.L().Error("parse-account-id-error", zap.Error(err))
		_ = writeResponse(w, []byte(``), http.StatusBadRequest)
		return
	}

	account, err := a.accountService.Find(accountID)
	if err != nil {
		zap.L().Error("account-find-error: %w", zap.Error(err))
		_ = writeResponse(w, []byte(``), http.StatusNotFound)
		return
	}

	accountResponse := payloads.AccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
	_ = writeResponse(w, accountResponse, http.StatusOK)
}

func NewAccountHandler(accountService service.AccountService) Account {
	return Account{
		accountService: accountService,
	}
}
