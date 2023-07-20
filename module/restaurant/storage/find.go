package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *mysqlStorage) GetRestaurant(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*model.Restaurant, error) {
	var data model.Restaurant

	db := s.DB

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Table(model.Restaurant{}.TableName()).Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
