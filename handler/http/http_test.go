package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"

	"transaction-api/handler/http/payloads"
)

type errWriterReader int

func (errWriterReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("reader-error")
}
func (errWriterReader) Write(_ []byte) (n int, err error) {
	return 0, errors.New("writer-error")
}
func (errWriterReader) Header() http.Header {
	return http.Header{}
}
func (errWriterReader) WriteHeader(_ int) {

}

func TestHttp(t *testing.T) {

	t.Run("Test Write Response 1 - Success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		response := payloads.AccountResponse{
			AccountID:      1,
			DocumentNumber: "1234567890",
		}

		assert.NilError(t, writeResponse(recorder, response, http.StatusOK))
		res := recorder.Result()
		assert.Equal(t, res.StatusCode, http.StatusOK)

	})

	t.Run("Test Write Response 2 - Error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		invalidResponse := map[string]interface{}{
			"foo": make(chan int),
		}

		err := writeResponse(rec, invalidResponse, http.StatusOK)
		assert.Error(t, err, "json: unsupported type: chan int")

	})

	t.Run("Test Write Response 3 - Error", func(t *testing.T) {
		var writer errWriterReader
		response := payloads.AccountResponse{
			AccountID:      1,
			DocumentNumber: "1234567890",
		}

		err := writeResponse(writer, response, http.StatusOK)
		assert.Error(t, err, "writer-error")

	})

	t.Run("Test Read Request 1 - Success", func(t *testing.T) {
		body := `{ "document_number": "1234567890" }`
		httpReq, err := http.NewRequest("", "", bytes.NewReader([]byte(body)))
		assert.NilError(t, err)

		request := payloads.AccountRequest{}
		assert.NilError(t, readRequest(httpReq, &request))
	})

	t.Run("Test Read Request 2 - Error", func(t *testing.T) {
		body := `{ "document_number": "" }`
		httpReq, err := http.NewRequest("", "", bytes.NewReader([]byte(body)))
		assert.NilError(t, err)

		request := payloads.AccountRequest{}
		assert.Error(t, readRequest(httpReq, &request), "document_number: cannot be blank.")
	})

	t.Run("Test Read Request 3 - Error", func(t *testing.T) {
		var reader errWriterReader
		httpReq, err := http.NewRequest("", "", reader)
		assert.NilError(t, err)

		request := payloads.AccountRequest{}
		assert.Error(t, readRequest(httpReq, &request), "reader-error")
	})

}
