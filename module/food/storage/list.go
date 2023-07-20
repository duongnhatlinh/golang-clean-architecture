package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

func (s *mysqlStorage) ListFood(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
	filter *model.Filter,
	moreKeys ...string,
) ([]model.Food, error) {
	var data []model.Food

	db := s.db.Table(model.Food{}.TableName()).Where(condition).Where("status in (1)")

	if f := filter; f != nil {
		if filter.Category_id > 0 {
			db = db.Where("category_id = ?", filter.Category_id)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", int(uid.GetLocalID()))
		}
	} else {
		db = db.Offset(paging.Limit * (paging.Page - 1))
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).
		Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return data, nil
}
