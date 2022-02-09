package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/model"
	"github.com/rafaelsanzio/go-core/pkg/repo"
	"github.com/rafaelsanzio/go-core/pkg/user"
)

func TestHandleGetUser(t *testing.T) {
	repo.SetUserRepo(repo.MockUserRepo{
		GetFunc: func(ctx context.Context, id string) (*user.User, errs.AppError) {
			if id == "1" {
				userMock := model.PrototypeUser()
				return &userMock, nil
			}

			return nil, nil
		},
	})
	defer repo.SetUserRepo(nil)

	testCases := []struct {
		Name               string
		ID                 string
		ExpectedStatusCode int
	}{
		{
			Name:               "Success handle get user",
			ID:                 "1",
			ExpectedStatusCode: 200,
		},
		{
			Name:               "Not Found handle get user",
			ID:                 "",
			ExpectedStatusCode: 404,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		req, err := http.NewRequest(http.MethodGet, "users/:id", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.ID})
		assert.NoError(t, err)
		res := httptest.NewRecorder()

		HandleGetUser(res, req)

		assert.Equal(t, tc.ExpectedStatusCode, res.Code)
		t.Logf("Response Body: %v", res.Body)

		if res.Code == 200 {
			user := user.User{}
			err = json.Unmarshal(res.Body.Bytes(), &user)
			assert.NoError(t, err)
		}
	}
}
