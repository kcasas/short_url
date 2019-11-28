package urlconv

import (
	"time"

	"github.com/enricofoltran/baseconv"
)

type URLShortener struct {
	db Storage
}

func New(db Storage) (*URLShortener, error) {
	return &URLShortener{
		db,
	}, nil
}

func (us *URLShortener) Shorten(longurl string, exp time.Duration) (shorturl string, err error) {
	id, err := us.db.InsertLongURL(longurl, exp)
	shorturl = baseconv.Base62.Encode(id)

	return shorturl, nil
}
