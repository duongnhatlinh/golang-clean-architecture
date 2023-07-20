package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/model"
	"gorm.io/gorm"
)

func (s *mysqlStorage) Delete(ctx context.Context, userId int, restaurantId int) error {
	db := s.db

	var like model.Restaurant_like

	if err := db.Table(model.Restaurant_like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDb(err)
	}

	if err := db.Table(model.Restaurant_like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
