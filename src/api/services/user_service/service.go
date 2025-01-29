package user_service

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/api/ports"
	"github.com/alaa-aqeel/todo/src/api/services/base_service"
	"github.com/alaa-aqeel/todo/src/api/validators"
	"github.com/alaa-aqeel/todo/src/helpers"
)

type UserService struct {
	Base *base_service.BaseService[models.User]
}

func (s UserService) Create(validated *validators.UserValidator) (*models.User, error) {
	user := &models.User{
		Name:     validated.Name,
		Username: validated.Username,
	}
	user.SetPassword(validated.Password)
	tx := s.Base.Db().Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (s UserService) Paginate(p ports.PaginatorPort, args helpers.Map) map[string]interface{} {

	return s.Base.Paginate(s.Filter(args), p).ToMap()
}
