package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

func (s *mysqlStorage) FindFood(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*model.Food, error) {
	var data model.Food

	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Table(model.Food{}.TableName()).Where(condition).First(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
