package base_service

import (
	"github.com/alaa-aqeel/todo/src/api/ports"
	"gorm.io/gorm"
)

func (b *BaseService[T]) ScopePaginate(p ports.PaginatorPort) ScopesCallback {

	return func(db *gorm.DB) *gorm.DB {

		return db.
			Offset(p.Offset()).
			Limit(p.Limit())
	}
}

func (b *BaseService[T]) Paginate(db *gorm.DB, p ports.PaginatorPort) ports.PaginatorPort {

	var model T
	p.SetTotal(b.CountOf(db, model))

	data := []*T{}
	db.Scopes(b.ScopePaginate(p)).Find(&data)
	p.SetData(data)

	return p
}
