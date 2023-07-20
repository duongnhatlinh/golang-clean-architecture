package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

func (s *mysqlStorage) GetListRestaurants(
	ctx context.Context,
	condition map[string]interface{},
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	var data []model.Restaurant

	db := s.DB

	db = db.Table(model.Restaurant{}.TableName()).Where(condition).Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if f := paging.FakeCursor; f != "" {
		if uid, err := common.FromBase58(f); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return data, nil
}
