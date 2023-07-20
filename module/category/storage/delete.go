package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

func (s *mysqlStorage) DeleteCategory(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.Category{}.TableName()).Where(condition).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
