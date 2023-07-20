package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

type CreateRestaurantStorage interface {
	CreateRestaurant(ctx context.Context, data *model.Restaurant) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStorage
}

func NewCreateRestaurantBiz(store CreateRestaurantStorage) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateNewRestaurant(ctx context.Context, data *model.Restaurant) error {
	err := biz.store.CreateRestaurant(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
