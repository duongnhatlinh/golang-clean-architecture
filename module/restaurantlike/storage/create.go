package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/model"
)

func (s *mysqlStorage) Create(ctx context.Context, data *model.Restaurant_like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
