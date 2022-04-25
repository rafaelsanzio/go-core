package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/repo"
)

func mockDeleteUserFunc(ctx context.Context, id string) errs.AppError {
	return nil
}

func mockDeleteUserThrowFunc(ctx context.Context, id string) errs.AppError {
	return errs.ErrRepoMockAction
}

func TestHandleDeleteUser(t *testing.T) {
	testCases := []struct {
		Name                 string
		ID                   string
		HandleDeleteFunction func(ctx context.Context, id string) errs.AppError
		ExpectedStatusCode   int
	}{
		{
			Name:                 "Success handle delete user",
			ID:                   "1",
			HandleDeleteFunction: mockDeleteUserFunc,
			ExpectedStatusCode:   204,
		},
		{
			Name:                 "Not Found handle delete user",
			ID:                   "",
			HandleDeleteFunction: mockDeleteUserFunc,
			ExpectedStatusCode:   404,
		},
		{
			Name:                 "Throwing error on function",
			ID:                   "1",
			HandleDeleteFunction: mockDeleteUserThrowFunc,
			ExpectedStatusCode:   500,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		repo.SetUserRepo(repo.MockUserRepo{
			DeleteFunc: tc.HandleDeleteFunction,
		})
		defer repo.SetUserRepo(nil)

		req, err := http.NewRequest(http.MethodDelete, "users/:id", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.ID})
		assert.NoError(t, err)
		res := httptest.NewRecorder()

		HandleDeleteUser(res, req)

		assert.Equal(t, tc.ExpectedStatusCode, res.Code)

		if res.Code == 204 {
			assert.NoError(t, err)
		}
	}
}
