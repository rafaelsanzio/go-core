package handlers

import (
	"bytes"
	"context"
	"encoding/json"
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

func TestHandlePostUser(t *testing.T) {
	goodReq := httptest.NewRequest(http.MethodPost, "/users", nil)
	goodReq = mux.SetURLVars(goodReq, map[string]string{})

	body, err := json.Marshal(UserEntityPayload{
		Name: "John Doe",
		Age:  "38",
	})
	assert.Equal(t, nil, err)

	goodReq.Body = ioutil.NopCloser(bytes.NewReader(body))

	noBodyReq := httptest.NewRequest(http.MethodPost, "/users", nil)
	noBodyReq = mux.SetURLVars(noBodyReq, map[string]string{})

	repo.SetUserRepo(repo.MockUserRepo{
		InsertFunc: func(ctx context.Context, user user.User) errs.AppError {
			return nil
		},
	})
	defer repo.SetUserRepo(nil)

	testCases := []struct {
		Name               string
		Request            *http.Request
		ExpectedStatusCode int
	}{
		{
			Name:               "Should return 201 if successful",
			Request:            goodReq,
			ExpectedStatusCode: 201,
		}, {
			Name:               "Should return 422 bad request",
			Request:            noBodyReq,
			ExpectedStatusCode: 422,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		w := httptest.NewRecorder()

		HandlePostUser(w, tc.Request)
		assert.Equal(t, tc.ExpectedStatusCode, w.Code)
	}
}

func TestConvertPayloadToUser(t *testing.T) {
	inPayload := UserEntityPayload{
		Name: "John Doe",
		Age:  "38",
	}

	expectedUser := user.User{
		ID:   "",
		Name: "John Doe",
		Age:  38,
	}

	testCases := []struct {
		Name          string
		Payload       UserEntityPayload
		ExpectedUser  user.User
		ExpectError   bool
		ExpectedError string
	}{
		{
			Name:         "Test Case: 1 - correct body, no error",
			Payload:      inPayload,
			ExpectedUser: expectedUser,
			ExpectError:  false,
		}, {
			Name:          "Test Case: 2 - error errs.ErrConvertingStringToInt",
			Payload:       UserEntityPayload{},
			ExpectError:   true,
			ExpectedError: "CMN010: error converting string to int, err: [strconv.Atoi: parsing \"\": invalid syntax]",
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		user, err := convertPayloadToUser(tc.Payload)
		if tc.ExpectError {
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), tc.ExpectedError)
		} else {
			assert.Equal(t, tc.ExpectedUser.ID, user.ID)
			assert.Equal(t, tc.ExpectedUser.Name, user.Name)
			assert.Equal(t, tc.ExpectedUser.Age, user.Age)
		}
	}
}

func TestDecodeUserRequest(t *testing.T) {
	goodReq := httptest.NewRequest(http.MethodPost, "/users", nil)
	goodReq = mux.SetURLVars(goodReq, map[string]string{})

	body, err := json.Marshal(UserEntityPayload{
		Name: "John Doe",
		Age:  "38",
	})
	assert.Equal(t, nil, err)

	goodReq.Body = ioutil.NopCloser(bytes.NewReader(body))

	noBodyReq := httptest.NewRequest(http.MethodPost, "/users", nil)
	noBodyReq = mux.SetURLVars(noBodyReq, map[string]string{})

	testCases := []struct {
		Name          string
		Request       *http.Request
		Payload       *UserEntityPayload
		ExpectedError bool
	}{
		{
			Name:    "Test Case: 1 - correct body, no error",
			Request: goodReq, Payload: &UserEntityPayload{
				Name: "John Doe",
				Age:  "38",
			}, ExpectedError: false,
		},
		{Name: "Test Case: 2 - no body, error found", Request: noBodyReq, Payload: nil, ExpectedError: true},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		decodedPayload, err := decodeUserRequest(tc.Request)
		if tc.ExpectedError {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, tc.Payload.Name, decodedPayload.Name)
			assert.Equal(t, tc.Payload.Age, decodedPayload.Age)
		}
	}
}
