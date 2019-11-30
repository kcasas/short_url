package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var defaultDB *gorm.DB

func init() {
	dsn := os.Getenv("DB_DSN")
	var err error
	defaultDB, err = New("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %w", err))
	}
}

func DB() *gorm.DB {
	return defaultDB
}

func Close() {
	defaultDB.Close()
}

func New(driver string, dsn string) (*gorm.DB, error) {
	database, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return database, nil
}
