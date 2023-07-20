package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

type GetRestaurantStorage interface {
	GetRestaurant(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStorage
}

func NewGetRestaurantBiz(store GetRestaurantStorage) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*model.Restaurant, error) {
	data, err := biz.store.GetRestaurant(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(model.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(model.EntityName, nil)
	}

	return data, nil
}
