package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/repo"
)

var jsonMarshal = json.Marshal

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		_ = errs.ErrGettingParam.Throwf(applog.Log, "error getting param, err: [%v]", id)
		httpNotFound(w)
		return
	}

	user, err := repo.GetUserRepo().Get(ctx, id)
	if err != nil {
		httpInternalServerError(w)
		return
	}

	data, err_ := jsonMarshal(user)
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
