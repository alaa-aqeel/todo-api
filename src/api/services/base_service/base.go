package base_service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewBaseService[T interface{}](db *gorm.DB) *BaseService[T] {
	return &BaseService[T]{
		db: db,
	}
}

func (s *BaseService[T]) Db() *gorm.DB {
	return s.db
}

func (s *BaseService[T]) Delete(pk string) bool {

	var model *T
	tx := s.db.Where("id = ?", pk).Delete(&model)

	return tx.Error == nil
}

func (s *BaseService[T]) Update(pk string, data map[string]interface{}) (records []T, tx *gorm.DB) {
	tx = s.db.Model(&records).
		Clauses(clause.Returning{}).
		Where("id = ?", pk).
		Updates(data)

	return
}
