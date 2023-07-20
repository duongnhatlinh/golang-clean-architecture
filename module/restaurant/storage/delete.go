package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

func (s *mysqlStorage) DeleteRestaurant(ctx context.Context, id int) error {
	db := s.DB

	if err := db.Table(model.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
