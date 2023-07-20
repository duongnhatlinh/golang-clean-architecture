package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

func (s *mysqlStorage) FindCategory(
	ctx context.Context,
	condition map[string]interface{},
) (*model.Category, error) {
	var data model.Category

	db := s.db

	if err := db.Table(model.Category{}.TableName()).Where(condition).First(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
