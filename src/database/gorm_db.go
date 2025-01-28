package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Orm *gorm.DB

func ConnectDatabase(databaseUrl string) *gorm.DB {
	var err error

	Orm, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return Orm
}
