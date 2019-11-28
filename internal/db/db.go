package db

import "github.com/jinzhu/gorm"

var defaultDB *gorm.DB

func init() {
	// dsn := os.Getenv("DB_DSN")
	// var err error
	// defaultDB, err = New("mysql", dsn)
	// if err != nil {
	// 	panic("mysql database connection error: " + dsn)
	// }
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
