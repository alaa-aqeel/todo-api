package models

import (
	"time"

	"github.com/alaa-aqeel/todo/src/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID          string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();not null"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null"`
	Username    string         `json:"username" gorm:"type:varchar(100);unique;not null"`
	Password    string         `json:"password" gorm:"type:varchar(255);not null"`
	IsActivated bool           `json:"is_activated" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (user *User) SetPassword(value string) (err error) {
	user.Password, err = helpers.MakeHash(value)

	return
}
