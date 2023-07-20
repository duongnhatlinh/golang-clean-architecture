package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/model"
)

func (s *mysqlStorage) Delete(ctx context.Context, userId int, foodId int) error {
	db := s.db

	if err := db.Table(model.Cart{}.TableName()).
		Where("user_id = ? and food_id = ?", userId, foodId).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
