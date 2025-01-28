package database

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}
