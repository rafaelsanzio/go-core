package handlers

import (
	"net/http"

	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
)

func httpUnprocessableEntity(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, err := w.Write([]byte(message))
	if err != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, "error writing response body, err: [%v]", err)
	}
}

func httpInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func httpNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
