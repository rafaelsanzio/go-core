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

	repo.SetUserRepo(repo.MockUserRepo{
		UpdateFunc: func(ctx context.Context, u user.User) (*user.User, errs.AppError) {
			return &u, nil
		},
	})
	defer repo.SetUserRepo(nil)

	testCases := []struct {
		Name               string
		Request            *http.Request
		ExpectedStatusCode int
	}{
		{
			Name:               "Should return 200 if successful",
			Request:            goodReq,
			ExpectedStatusCode: 200,
		}, {
			Name:               "Should return 422 bad request",
			Request:            noBodyReq,
			ExpectedStatusCode: 422,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		w := httptest.NewRecorder()

		HandleUpdateUser(w, tc.Request)
		assert.Equal(t, tc.ExpectedStatusCode, w.Code)
	}
}
