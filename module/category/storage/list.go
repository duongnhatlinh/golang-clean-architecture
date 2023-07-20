package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

func (s *mysqlStorage) ListCategories(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
) ([]model.Category, error) {
	var data []model.Category

	db := s.db.Table(model.Category{}.TableName()).Where(condition).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
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
