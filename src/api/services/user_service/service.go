package user_service

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/api/validators"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (s UserService) Create(validated *validators.UserValidator) (*models.User, error) {

	user := &models.User{
		Name:     validated.Name,
		Username: validated.Username,
	}
	user.SetPassword(validated.Password)

	tx := s.Db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
