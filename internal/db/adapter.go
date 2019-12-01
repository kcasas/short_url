package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kcasas/short_url/internal/db/models"
)

type Adapter struct {
	db *gorm.DB
}

type DBAdapter interface {
	CreateID() (int64, error)
	SaveURL(short string, longURL string, exp int64) error
	ExpandURL(short string) (longURL string, err error)
}

func NewAdapter(db *gorm.DB) *Adapter {
	return &Adapter{db}
}

func (a *Adapter) CreateID() (int64, error) {
	i := models.ID{
		ID: 0,
	}
	err := a.db.Create(&i).Error

	return i.ID, err
}

func (a *Adapter) SaveURL(short string, long string, expSecs int64) error {
	urlModel := models.URL{
		Short:     short,
		LongURL:   long,
		CreatedAt: time.Now().UTC(),
	}

	switch {
	case expSecs == -1:
	case expSecs > 0:
		expAt := urlModel.CreatedAt.Add(
			time.Duration(expSecs) * time.Second,
		)
		urlModel.ExpAt = &expAt
	default:
		expAt := urlModel.CreatedAt.Add(
			time.Duration(24 * time.Hour),
		)
		urlModel.ExpAt = &expAt
	}

	return a.db.Create(&urlModel).Error
}

func (a *Adapter) ExpandURL(short string) (longURL string, err error) {
	urlModel := models.URL{}

	err = a.db.Raw(
		"SELECT long_url FROM urls WHERE short = ? AND (exp_at > ? OR exp_at IS NULL) LIMIT 1",
		short,
		time.Now().UTC(),
	).Scan(&urlModel).Error

	return urlModel.LongURL, err
}
