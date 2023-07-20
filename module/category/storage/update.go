package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

func (s *mysqlStorage) UpdateCategory(
	ctx context.Context,
	condition map[string]interface{},
	data *model.UpdateCategory,
) error {
	db := s.db

	if err := db.Where(condition).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
