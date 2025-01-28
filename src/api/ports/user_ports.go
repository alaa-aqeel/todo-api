package ports

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/api/validators"
)

type UserServicePort interface {
	Create(user *validators.UserValidator) (*models.User, error)
}
