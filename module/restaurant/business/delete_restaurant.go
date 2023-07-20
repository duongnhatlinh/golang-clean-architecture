package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

type DeleteRestaurantStorage interface {
	GetRestaurant(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
	DeleteRestaurant(ctx context.Context, id int) error
}

type deleteRestaurantsBiz struct {
	store DeleteRestaurantStorage
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStorage) *deleteRestaurantsBiz {
	return &deleteRestaurantsBiz{store: store}
}

func (biz *deleteRestaurantsBiz) DeleteRestaurant(ctx context.Context, id int) error {
	data, err := biz.store.GetRestaurant(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, nil)
	}

	if err := biz.store.DeleteRestaurant(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
