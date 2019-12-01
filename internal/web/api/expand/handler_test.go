package expand_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kcasas/short_url/internal/db/mocks"
	"github.com/kcasas/short_url/internal/web/api/expand"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	r := require.New(t)

	short := "foo"
	jsonPayload := fmt.Sprintf(`{"short":"%s"}`, short)

	long := "foobar"
	jsonResponse := fmt.Sprintf(`{"long":"%s"}`, long)
	testcases := []struct {
		name       string
		err        error
		statusCode int
	}{
		{"no error", nil, http.StatusOK},
		{"has error", errors.New("Not Found"), http.StatusNotFound},
	}

	for _, tc := range testcases {
		req, err := http.NewRequest(
			http.MethodPost, "/",
			strings.NewReader(jsonPayload))
		if err != nil {
			r.NoError(err)
		}

		dbAdapter := mocks.DBAdapter{}

		dbAdapter.On("ExpandURL", short).Return(long, tc.err)
		handler := expand.NewExpandHandler(&dbAdapter)
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, req)

		r.Equal(tc.statusCode, resp.Code)
		if tc.err == nil {
			r.JSONEq(jsonResponse, resp.Body.String())
		}
	}
}
