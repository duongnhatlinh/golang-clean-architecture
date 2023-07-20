package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/model"
)

func (s *mysqlStorage) CreateUser(ctx context.Context, data *model.User) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
