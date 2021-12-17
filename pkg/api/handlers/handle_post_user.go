package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/repo"
	"github.com/rafaelsanzio/go-core/pkg/user"
)

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
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

	err = repo.GetUserRepo().Insert(ctx, user)
	if err != nil {
		httpInternalServerError(w)
		return
	}

	w.WriteHeader(200)
}

func decodeUserRequest(r *http.Request) (UserEntityPayload, errs.AppError) {
	payload := UserEntityPayload{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		return payload, errs.ErrUnmarshalingJson.Throwf(applog.Log, errs.ErrFmt, err)
	}

	return payload, nil
}

func convertPayloadToUser(u UserEntityPayload) (user.User, errs.AppError) {
	age, err := strconv.Atoi(u.Age)
	if err != nil {
		return user.User{}, errs.ErrConvertingStringToInt.Throwf(applog.Log, errs.ErrFmt, err)
	}
	result := user.User{
		Name: u.Name,
		Age:  age,
	}

	return result, nil
}
