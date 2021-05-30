package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	dbURL := "ovunque:qazxc1234@tcp(127.0.0.1:3306)/gin_getting_started"
	db, err := gorm.Open(mysql.Open(dbURL))

	if err != nil {
		log.Fatalf("Error connection to DB, %v", err)
	}

	return db
}
