package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/repo"
	"github.com/rafaelsanzio/go-core/pkg/user"
)

func mockUpdateUserFunc(ctx context.Context, u user.User) (*user.User, errs.AppError) {
	return &u, nil
}

func mockUpdateUserThrowFunc(ctx context.Context, u user.User) (*user.User, errs.AppError) {
	return nil, errs.ErrRepoMockAction
}

func TestHandleUpdateUser(t *testing.T) {
	goodReq := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/users/%s", "1"), nil)
	goodReq = mux.SetURLVars(goodReq, map[string]string{})

	body, err := json.Marshal(UserEntityPayload{
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "john@mail.com",
	})
	assert.Equal(t, nil, err)

	goodReq.Body = ioutil.NopCloser(bytes.NewReader(body))

	noBodyReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", "1"), nil)
	noBodyReq = mux.SetURLVars(noBodyReq, map[string]string{})

	throwReq := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/users/%s", "1"), nil)
	throwReq = mux.SetURLVars(throwReq, map[string]string{})

	throwReq.Body = ioutil.NopCloser(bytes.NewReader(body))

	testCases := []struct {
		Name                 string
		Request              *http.Request
		HandleUpdateFunction func(ctx context.Context, u user.User) (*user.User, errs.AppError)
		ExpectedStatusCode   int
	}{
		{
			Name:                 "Should return 200 if successful",
			Request:              goodReq,
			HandleUpdateFunction: mockUpdateUserFunc,
			ExpectedStatusCode:   200,
		}, {
			Name:                 "Throwing error on function",
			Request:              throwReq,
			HandleUpdateFunction: mockUpdateUserThrowFunc,
			ExpectedStatusCode:   500,
		}, {
			Name:                 "Should return 422 bad request",
			Request:              noBodyReq,
			HandleUpdateFunction: mockUpdateUserFunc,
			ExpectedStatusCode:   422,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		repo.SetUserRepo(repo.MockUserRepo{
			UpdateFunc: tc.HandleUpdateFunction,
		})
		defer repo.SetUserRepo(nil)

		w := httptest.NewRecorder()

		HandleUpdateUser(w, tc.Request)
		fmt.Println("w.Code", w.Code)
		assert.Equal(t, tc.ExpectedStatusCode, w.Code)
	}
}
