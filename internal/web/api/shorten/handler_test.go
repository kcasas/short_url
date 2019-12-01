package shorten_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/enricofoltran/baseconv"
	urlconv_mocks "github.com/kcasas/short_url/internal/urlconv/mocks"
	"github.com/kcasas/short_url/internal/web/api/shorten"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	r := require.New(t)

	testcases := []struct {
		name       string
		url        string
		err        error
		statusCode int
	}{
		{"no error", "https://example.com", nil, http.StatusOK},
		{"bad url format", "", nil, http.StatusBadRequest},
		{"save error", "https://example.com", errors.New("save error"), http.StatusInternalServerError},
	}
	for _, tc := range testcases {
		exp := int64(-1)
		idPrefix := int64(30)
		id := int64(20)

		jsonPayload := fmt.Sprintf(
			`{"url":"%s","expiration":%d}`,
			tc.url,
			int64(-1),
		)
		req, err := http.NewRequest(
			http.MethodPost, "/",
			strings.NewReader(jsonPayload))
		if err != nil {
			r.NoError(err)
		}

		prefixer := urlconv_mocks.Prefixer{}
		prefixer.On("CreateIDPrefix").Return(int(idPrefix), nil)
		dbAdapter := urlconv_mocks.DBAdapter{}
		dbAdapter.On("CreateID").Return(id, nil)

		short := baseconv.Base62.Encode(idPrefix) + baseconv.Base62.Encode(id)
		dbAdapter.On("SaveURL", short, tc.url, exp).Return(tc.err)

		handler := shorten.NewShortenHandler(&dbAdapter, &prefixer)
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, req)

		r.Equal(tc.statusCode, resp.Code)
		if resp.Code < http.StatusBadRequest {
			jsonResponse := shorten.JsonResponse{}
			err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
			r.NoError(err)
			r.Equal(tc.url, jsonResponse.Long)
			r.NotEmpty(jsonResponse.Short)
		}
	}
}
