package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
)

func writeResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data, err := json.Marshal(body)
	if err != nil {
		zap.L().Error("json-marshal-error", zap.Error(err))
		return
	}

	if _, err := w.Write(data); err != nil {
		zap.L().Error("http-response-writer-error", zap.Error(err))
	}
}

func readRequest(r *http.Request, into interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		zap.L().Error("request-body-read-error", zap.Error(err))
		return err
	}

	if err := json.Unmarshal(body, &into); err != nil {
		zap.L().Error("json-unmarshal-error", zap.Error(err))
		return err
	}

	if val, ok := into.(validation.Validatable); ok {
		if err = val.Validate(); err != nil {
			zap.L().Error("data-validation-error", zap.Error(err))
			return err
		}
	}

	return nil
}
