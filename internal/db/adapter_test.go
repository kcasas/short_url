package db_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/kcasas/short_url/internal/db"
	"github.com/kcasas/short_url/internal/db/models"
	"github.com/stretchr/testify/require"
)

func TestCreateID(t *testing.T) {
	r := require.New(t)
	a := db.NewAdapter(db.DB())
	ids := map[int64]bool{}
	for i := 0; i < 1000; i++ {
		id, err := a.CreateID()
		r.NoError(err)
		_, exists := ids[id]
		r.False(exists, "create id can produce duplicates")
		ids[id] = false
	}
}

func TestSaveURLOnDuplicateShort(t *testing.T) {
	r := require.New(t)
	a := db.NewAdapter(db.DB())
	id, err := a.CreateID()
	r.NoError(err)
	short := strconv.Itoa(int(id))

	err = a.SaveURL(short, "", int64(-1))
	r.NoError(err)
	err = a.SaveURL(short, "", int64(-1))
	r.Error(err)
}

func TestSaveURLExpiration(t *testing.T) {
	r := require.New(t)
	a := db.NewAdapter(db.DB())

	testcases := []struct {
		name       string
		expiration int64
		expected   time.Duration
	}{
		{
			name:       "-1 means no expiration",
			expiration: -1,
		},
		{
			name:     "unset means 24 hours",
			expected: time.Duration(24) * time.Hour,
		},
		{
			name:       "expiration in 60 seconds",
			expiration: 60,
			expected:   time.Duration(60) * time.Second,
		},
		{
			name:       "expiration in 3000 seconds",
			expiration: 3000,
			expected:   time.Duration(3000) * time.Second,
		},
	}

	for _, tc := range testcases {
		id, err := a.CreateID()
		r.NoError(err)
		short := strconv.Itoa(int(id))
		err = a.SaveURL(short, "", tc.expiration)
		r.NoError(err)

		urlModel := models.URL{}
		err = db.DB().Raw(
			"SELECT created_at, exp_at FROM urls WHERE short = ? AND (exp_at > ? OR exp_at IS NULL) LIMIT 1",
			short,
			time.Now().UTC(),
		).Scan(&urlModel).Error
		r.NoError(err)

		var diff time.Duration
		if urlModel.ExpAt != nil {
			diff = urlModel.ExpAt.Sub(urlModel.CreatedAt)
		}

		r.Equal(tc.expected, diff)
	}
}

func TestExpandURL(t *testing.T) {
	r := require.New(t)
	a := db.NewAdapter(db.DB())
	longurl := "https://example.com"

	id, err := a.CreateID()
	r.NoError(err)
	short := strconv.Itoa(int(id))
	err = a.SaveURL(short, longurl, int64(-1))
	r.NoError(err)

	actual, err := a.ExpandURL(short)
	r.NoError(err)
	r.Equal(longurl, actual)
}
