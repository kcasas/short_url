package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kcasas/short_url/internal/db/models"
)

type Adapter struct {
	db *gorm.DB
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
