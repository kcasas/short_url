package urlconv_test

import (
	"testing"

	"github.com/enricofoltran/baseconv"
	"github.com/kcasas/short_url/internal/urlconv"
	"github.com/kcasas/short_url/internal/urlconv/mocks"
	"github.com/stretchr/testify/require"
)

func TestShorten(t *testing.T) {
	r := require.New(t)
	exp := int64(-1)
	idPrefix := int64(30)
	id := int64(20)
	longURL := "https://example.com"

	prefixer := mocks.Prefixer{}
	prefixer.On("CreateIDPrefix").Return(int(idPrefix), nil)
	dbAdapter := mocks.DBAdapter{}
	dbAdapter.On("CreateID").Return(id, nil)

	short := baseconv.Base62.Encode(idPrefix) + baseconv.Base62.Encode(id)
	dbAdapter.On("SaveURL", short, longURL, exp).Return(nil)

	shortener := urlconv.NewShortener(&dbAdapter, &prefixer)
	actual, err := shortener.Shorten(longURL, exp)
	r.NoError(err)
	r.Equal(short, actual)
}
