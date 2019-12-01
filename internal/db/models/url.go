package models

import (
	"time"
)

type URL struct {
	Short     string     `gorm:"primary_key"`
	LongURL   string     `gorm:"column:long_url"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	ExpAt     *time.Time `gorm:"column:exp_at"`
}
