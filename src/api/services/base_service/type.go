package base_service

import (
	"gorm.io/gorm"
)

type ScopesCallback func(db *gorm.DB) *gorm.DB

type BaseService[T interface{}] struct {
	db *gorm.DB
}
