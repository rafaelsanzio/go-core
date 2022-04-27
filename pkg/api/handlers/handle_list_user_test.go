package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/model"
	"github.com/rafaelsanzio/go-core/pkg/repo"
	"github.com/rafaelsanzio/go-core/pkg/user"
)

func mockListUserFunc(ctx context.Context) ([]user.User, errs.AppError) {
	userMock := model.PrototypeUser()

	userMock2 := model.PrototypeUser()
	userMock2.FirstName = "John 2"

	userMockList := []user.User{userMock, userMock2}

	return userMockList, nil
}

func mockListUserThrowFunc(ctx context.Context) ([]user.User, errs.AppError) {
	return nil, errs.ErrRepoMockAction
}

func TestHandleListUser(t *testing.T) {
	testCases := []struct {
		Name                   string
		HandleListUserFunction func(ctx context.Context) ([]user.User, errs.AppError)
		ExpectedStatusCode     int
	}{
		{
			Name:                   "Success handle list user",
			HandleListUserFunction: mockListUserFunc,
			ExpectedStatusCode:     200,
		},
		{
			Name:                   "Throwing handle list user",
			HandleListUserFunction: mockListUserThrowFunc,
			ExpectedStatusCode:     500,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		repo.SetUserRepo(repo.MockUserRepo{
			ListFunc: tc.HandleListUserFunction,
		})
		defer repo.SetUserRepo(nil)

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()

		HandleListUser(res, req)

		assert.Equal(t, tc.ExpectedStatusCode, res.Code)
		t.Logf("Response Body: %v", res.Body)

		if res.Code == 200 {
			user := []user.User{}
			err = json.Unmarshal(res.Body.Bytes(), &user)
			assert.NoError(t, err)

			assert.Equal(t, 2, len(user))
		}
	}
}
