package handler

import "net/http"

type AccountHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type TransactionHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
}
