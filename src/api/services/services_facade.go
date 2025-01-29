package services

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/api/ports"
	"github.com/alaa-aqeel/todo/src/api/services/base_service"
	"github.com/alaa-aqeel/todo/src/api/services/user_service"
	"github.com/alaa-aqeel/todo/src/database"
)

func UserService() ports.UserServicePort {

	return &user_service.UserService{
		Base: base_service.NewBaseService[models.User](database.Orm),
	}
}
