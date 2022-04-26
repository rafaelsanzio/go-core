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

func mockGetUserFunc(ctx context.Context, id string) (*user.User, errs.AppError) {
	if id == "1" {
		userMock := model.PrototypeUser()
		return &userMock, nil
	}
	return nil, nil
}

func mockGetUserThrowFunc(ctx context.Context, id string) (*user.User, errs.AppError) {
	return nil, errs.ErrRepoMockAction
}

func TestHandleGetUser(t *testing.T) {
	testCases := []struct {
		Name                  string
		ID                    string
		HandleGetUserFunction func(ctx context.Context, id string) (*user.User, errs.AppError)
		ExpectedStatusCode    int
	}{
		{
			Name:                  "Success handle get user",
			ID:                    "1",
			HandleGetUserFunction: mockGetUserFunc,
			ExpectedStatusCode:    200,
		},
		{
			Name:                  "Not Found handle get user",
			ID:                    "",
			HandleGetUserFunction: mockGetUserFunc,
			ExpectedStatusCode:    404,
		},
		{
			Name:                  "Error getting repo user",
			ID:                    "1",
			HandleGetUserFunction: mockGetUserThrowFunc,
			ExpectedStatusCode:    500,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		repo.SetUserRepo(repo.MockUserRepo{
			GetFunc: tc.HandleGetUserFunction,
		})
		defer repo.SetUserRepo(nil)

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
