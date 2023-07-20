package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/model"
)

func (s *mysqlStorage) FindFood(
	ctx context.Context,
	userId int,
	foodId int,
) (*model.Cart, error) {
	var data model.Cart

	db := s.db

	if err := db.Table(model.Cart{}.TableName()).Where("user_id = ? and food_id = ?", userId, foodId).First(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
