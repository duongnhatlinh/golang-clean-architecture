package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

func (s *mysqlStorage) CreateRestaurant(ctx context.Context, data *model.Restaurant) error {

	if err := s.DB.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
