package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

func (s *mysqlStorage) CreateCategory(ctx context.Context, data *model.Category) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
