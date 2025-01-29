package ports

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/api/validators"
	"github.com/alaa-aqeel/todo/src/helpers"
)

type UserServicePort interface {
	Create(user *validators.UserValidator) (*models.User, error)
	Paginate(p PaginatorPort, filters helpers.Map) map[string]interface{}
	Get(pk string) *models.User
}
