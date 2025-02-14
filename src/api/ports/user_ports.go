package ports

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/helpers"
)

type UserServicePort interface {
	Create(data map[string]interface{}) (*models.User, error)
	Paginate(p PaginatorPort, filters helpers.Map) map[string]interface{}
	Get(pk string) *models.User
	Delete(pk string) bool
	Update(pk string, data map[string]interface{}) ([]models.User, error)
}
