package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

func (s *mysqlStorage) DeleteFood(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.Food{}.TableName()).Where(condition).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
