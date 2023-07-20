package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/model"
)

func (s *mysqlStorage) List(ctx context.Context, userId int, moreKeys ...string) ([]model.Cart, error) {
	var data []model.Cart

	db := s.db.Table(model.Cart{}.TableName()).Where("status in (1)")

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where("user_id = ?", userId).Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return data, nil
}
