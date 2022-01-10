package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/repo"
)

func HandleListUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := repo.GetUserRepo().List(ctx)
	if err != nil {
		httpInternalServerError(w)
		return
	}

	data, err_ := json.Marshal(user)
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

	w.WriteHeader(200)
}
