package user_service

import (
	"github.com/alaa-aqeel/todo/src/api/models"
	"github.com/alaa-aqeel/todo/src/api/ports"
	"github.com/alaa-aqeel/todo/src/helpers"
	"gorm.io/gorm"
)

func (s UserService) Filter(args helpers.Map) (query *gorm.DB) {
	query = s.Base.Db()
	if args.Has("id") {
		query = query.Where("id = ?", args.GetString("id"))
	}
	if args.Has("username") {
		query = query.Where("username = ?", args.GetString("username"))
	}
	if args.Has("name") {
		query = query.Where("name LIKE ?", "%"+args.GetString("name")+"%")
	}

	return
}

func (s UserService) Get(pk string) (user *models.User) {

	s.Filter(helpers.Map{"id": pk}).First(&user)

	return
}

func (s UserService) Paginate(p ports.PaginatorPort, args helpers.Map) map[string]interface{} {

	return s.Base.Paginate(s.Filter(args), p).ToMap()
}
