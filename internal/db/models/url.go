package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type URL struct {
	Short     string     `gorm:"primary_key"`
	LongURL   string     `gorm:"column:long_url"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	ExpAt     *time.Time `gorm:"column:exp_at"`
}

func (u *URL) Expand(DB *gorm.DB) error {
	return DB.Raw(
		"SELECT long_url FROM urls WHERE short = ? AND (exp_at > ? OR exp_at IS NULL) LIMIT 1",
		u.Short,
		time.Now().UTC(),
	).Scan(&u).Error
}
