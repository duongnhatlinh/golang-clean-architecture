package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/model"
)

func (s *mysqlStorage) UpdateUser(ctx context.Context, id int, data *model.UserUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
