package http

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
	"transaction-api/entity"
	"transaction-api/handler/http/payloads"
	"transaction-api/repository"
)

type Account struct {
	accountRepo repository.AccountRepository
}

func (a Account) Post(w http.ResponseWriter, r *http.Request) {
	accountRequest := payloads.AccountRequest{}
	if err := readRequest(r, &accountRequest); err != nil {
		writeResponse(w, []byte(``), http.StatusBadRequest)
		return
	}

	now := time.Now()
	account := entity.Account{
		DocumentNumber: accountRequest.DocumentNumber,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := a.accountRepo.Create(&account); err != nil {
		zap.L().Error("account-store-error", zap.Error(err))
		writeResponse(w, []byte(``), http.StatusInternalServerError)
		return
	}
	zap.L().Info("account-created", zap.Any("account", account))

	accountResponse := payloads.AccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
	writeResponse(w, accountResponse, http.StatusOK)
}

func (a Account) Get(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		zap.L().Error("parse-account-id-error", zap.Error(err))
		writeResponse(w, []byte(``), http.StatusBadRequest)
		return
	}

	account, err := a.accountRepo.Find(int32(accountID))
	if err != nil {
		zap.L().Error("account-find-error: %w", zap.Error(err))
		writeResponse(w, []byte(``), http.StatusNotFound)
		return
	}

	accountResponse := payloads.AccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
	writeResponse(w, accountResponse, http.StatusOK)
}

func NewAccountHandler(accountRepo repository.AccountRepository) Account {
	return Account{
		accountRepo: accountRepo,
	}
}
