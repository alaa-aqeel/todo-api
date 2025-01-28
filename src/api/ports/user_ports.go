package ports

import "github.com/alaa-aqeel/todo/src/api/models"

type UserServicePort interface {
	All() []models.User
	Find(id string) models.User
	Create(user models.User) models.User
	Update(id string, user models.User) models.User
	Delete(id string) bool
}
