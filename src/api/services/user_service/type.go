package user_service

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/api/services/base_service"
)

type UserService struct {
	Base *base_service.BaseService[models.User]
}
