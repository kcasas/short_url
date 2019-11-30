package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var defaultDB *gorm.DB

func init() {
	dsn := viper.GetString("DB_DSN")
	if len(dsn) <= 0 {
		logrus.Fatal(errors.New("DB_DSN is empty"))
	}

	var err error
	defaultDB, err = New("mysql", dsn)
	if err != nil {
		logrus.Fatal(fmt.Errorf("unable to connect to database: %w", err))
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
