package db

import (
	"future-admin/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	if !env.IsProduction() {
		db = db.Debug()
	}
	return db, nil
}
