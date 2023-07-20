package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/foodlike/model"
)

func (s *mysqlStorage) Create(ctx context.Context, data *model.Food_like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
