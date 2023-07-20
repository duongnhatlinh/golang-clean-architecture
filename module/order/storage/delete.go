package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/model"
)

func (s *mysqlStorage) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.Order{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
