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

func TestHandleDeleteUser(t *testing.T) {
	repo.SetUserRepo(repo.MockUserRepo{
		DeleteFunc: func(ctx context.Context, id string) errs.AppError {
			if id == "1" {
				return nil
			}

			return nil
		},
	})
	defer repo.SetUserRepo(nil)

	testCases := []struct {
		Name               string
		ID                 string
		ExpectedStatusCode int
	}{
		{
			Name:               "Success handle delete user",
			ID:                 "1",
			ExpectedStatusCode: 204,
		},
		{
			Name:               "Not Found handle delete user",
			ID:                 "",
			ExpectedStatusCode: 404,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

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
