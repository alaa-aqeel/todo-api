package base_service

import "gorm.io/gorm"

func NewBaseService[T interface{}](db *gorm.DB) *BaseService[T] {
	return &BaseService[T]{
		db: db,
	}
}

func (s *BaseService[T]) Db() *gorm.DB {
	return s.db
}

func (s *BaseService[T]) CountOf(db *gorm.DB, model interface{}) int64 {
	var total int64
	db.Model(model).Count(&total)

	return total
}

func (s *BaseService[T]) Get(pk string) *T {

	var model *T
	s.db.Where("id = ?", pk).First(&model)
	return model
}
