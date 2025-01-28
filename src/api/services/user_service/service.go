package user_service

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (s UserService) Create(validated models.User) models.User {

	s.Db.Create(&validated)

	return validated
}
