package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/repo"
)

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		_ = errs.ErrGettingParam.Throwf(applog.Log, "error getting param, err: [%v]", id)
		httpNotFound(w)
		return
	}

	err := repo.GetUserRepo().Delete(ctx, id)
	if err != nil {
		httpInternalServerError(w)
		return
	}

	w.WriteHeader(204)
}
