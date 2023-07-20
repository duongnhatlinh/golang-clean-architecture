package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

func (s *mysqlStorage) CreateFood(ctx context.Context, data *model.Food) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}