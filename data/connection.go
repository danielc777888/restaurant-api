package data

import (
	"middleearth/eateries/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Get database connection.
func Connection() *gorm.DB {
	var dsn = env.DbDsn()
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
