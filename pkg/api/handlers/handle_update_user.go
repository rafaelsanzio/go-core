package handlers

import (
	"net/http"

	"github.com/rafaelsanzio/go-core/pkg/repo"
)

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userPayload, err := decodeUserRequest(r)
	if err != nil {
		httpUnprocessableEntity(w, err.Error())
		return
	}

	user, err := convertPayloadToUser(userPayload)
	if err != nil {
		httpUnprocessableEntity(w, err.Error())
	}

	_, err = repo.GetUserRepo().Update(ctx, user)
	if err != nil {
		httpInternalServerError(w)
		return
	}

	w.WriteHeader(200)
}
