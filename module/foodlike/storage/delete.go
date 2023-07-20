package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/foodlike/model"
	"gorm.io/gorm"
)

func (s *mysqlStorage) Delete(ctx context.Context, userId int, foodId int) error {
	db := s.db

	var like model.Food_like

	if err := db.Table(model.Food_like{}.TableName()).
		Where("user_id = ? and food_id = ?", userId, foodId).
		First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDb(err)
	}

	if err := db.Table(model.Food_like{}.TableName()).
		Where("user_id = ? and food_id = ?", userId, foodId).
		Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
