package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandleOk(t *testing.T) {
	goodReq := httptest.NewRequest(http.MethodGet, "/ok", nil)
	goodReq = mux.SetURLVars(goodReq, map[string]string{})

	testCases := []struct {
		Name               string
		Request            *http.Request
		ExpectedStatusCode int
	}{
		{
			Name:               "Should return 200 if successful",
			Request:            goodReq,
			ExpectedStatusCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		w := httptest.NewRecorder()

		HandleOK(w, tc.Request)
		assert.Equal(t, tc.ExpectedStatusCode, w.Code)
	}
}
