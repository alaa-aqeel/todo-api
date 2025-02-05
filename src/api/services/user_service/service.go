package user_service

import (
	"github.com/alaa-aqeel/todo/src/api/models"
)

func (s UserService) Create(validated map[string]interface{}) (user *models.User, err error) {

	err = s.Base.Db().Model(&user).Create(validated).Error
	return
}

func (s UserService) Update(pk string, data map[string]interface{}) ([]models.User, error) {

	records, tx := s.Base.Update(pk, data)
	return records, tx.Error
}

func (s UserService) Delete(pk string) bool {

	return s.Base.Delete(pk)
}
