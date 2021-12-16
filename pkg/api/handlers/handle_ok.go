package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
)

func HandleOK(w http.ResponseWriter, r *http.Request) {
	dataReturn := OkPayload{
		Health: 1,
		Test:   "Everthing is OK",
	}

	data, err_ := json.Marshal(dataReturn)
	if err_ != nil {
		_ = errs.ErrMarshalingJson.Throwf(applog.Log, errs.ErrFmt, err_)
		httpInternalServerError(w)
		return
	}

	_, err_ = w.Write(data)

	if err_ != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, "error writing timeline to response, err: [%v]", err_)
		httpInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
