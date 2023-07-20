package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/model"
	"gorm.io/gorm"
)

func (s *mysqlStorage) FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*model.User, error) {
	var user model.User

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Table(model.User{}.TableName()).Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDb(err)
	}

	return &user, nil
}
