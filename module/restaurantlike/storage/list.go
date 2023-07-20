package storage

import (
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/model"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *mysqlStorage) List(
	ctx context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	filter *model.Filter,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var data []model.Restaurant_like

	db := s.db.Table(model.Restaurant_like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDb(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("created_at desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	users := make([]common.SimpleUser, len(data))

	for i, item := range data {
		data[i].User.CreatedAt = item.CreatedAt
		data[i].User.UpdatedAt = nil
		users[i] = *data[i].User

		if i == len(data)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
