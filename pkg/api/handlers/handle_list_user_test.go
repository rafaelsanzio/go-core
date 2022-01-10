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

func TestHandleListUser(t *testing.T) {
	repo.SetUserRepo(repo.MockUserRepo{
		ListFunc: func(ctx context.Context) ([]user.User, errs.AppError) {
			userMock := model.PrototypeUser()

			userMock2 := model.PrototypeUser()
			userMock2.Name = "John Doe 2"

			userMockList := []user.User{userMock, userMock2}

			return userMockList, nil
		},
	})
	defer repo.SetUserRepo(nil)

	testCases := []struct {
		name               string
		expectedStatusCode int
	}{
		{
			"Success handle list user",
			200,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		req, err := http.NewRequest(http.MethodGet, "users", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()

		HandleListUser(res, req)

		assert.Equal(t, tc.expectedStatusCode, res.Code)
		t.Logf("Response Body: %v", res.Body)

		if res.Code == 200 {
			user := []user.User{}
			err = json.Unmarshal(res.Body.Bytes(), &user)
			assert.NoError(t, err)

			assert.Equal(t, 2, len(user))
		}
	}
}
