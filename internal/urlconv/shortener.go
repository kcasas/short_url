package urlconv

import (
	"fmt"

	"github.com/enricofoltran/baseconv"
)

type DBAdapter interface {
	CreateID() (int64, error)
	SaveURL(short string, long string, exp int64) error
}

type URLShortener struct {
	db       DBAdapter
	prefixer Prefixer
}

func NewShortener(adapter DBAdapter, prefixer Prefixer) *URLShortener {
	return &URLShortener{adapter, prefixer}
}

func (us *URLShortener) getPrefix() (string, error) {
	if us.prefixer == nil {
		return "", nil
	}

	prefixID, err := us.prefixer.CreateIDPrefix()
	if err != nil {
		return "", err
	}

	return baseconv.Base62.Encode(int64(prefixID)), nil
}

func (us *URLShortener) Shorten(longurl string, expSecs int64) (string, error) {
	prefix, err := us.getPrefix()
	if err != nil {
		return "", nil
	}

	id, err := us.db.CreateID()
	if err != nil {
		return "", err
	}

	shorturl := fmt.Sprintf(
		"%s%s",
		prefix,
		baseconv.Base62.Encode(id),
	)

	return shorturl, us.db.SaveURL(shorturl, longurl, expSecs)
}
