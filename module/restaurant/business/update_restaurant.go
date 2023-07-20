package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

type UpdateRestaurantStorage interface {
	GetRestaurant(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
	UpdateRestaurant(ctx context.Context, id int, data *model.UpdateRestaurant) error
}

type updateRestaurantsBiz struct {
	store UpdateRestaurantStorage
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStorage) *updateRestaurantsBiz {
	return &updateRestaurantsBiz{store: store}
}

func (biz *updateRestaurantsBiz) UpdateRestaurant(ctx context.Context, id int, data *model.UpdateRestaurant) error {
	dataOld, err := biz.store.GetRestaurant(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if dataOld.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, nil)
	}

	if err := biz.store.UpdateRestaurant(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
