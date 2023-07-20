package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
)

type GetListRestaurantsStorage interface {
	GetListRestaurants(
		ctx context.Context,
		condition map[string]interface{},
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Restaurant, error)
}

type getListRestaurantsBiz struct {
	store GetListRestaurantsStorage
}

func NewGetListRestaurantBiz(store GetListRestaurantsStorage) *getListRestaurantsBiz {
	return &getListRestaurantsBiz{store: store}
}

func (biz *getListRestaurantsBiz) ListRestaurants(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	data, err := biz.store.GetListRestaurants(ctx, nil, filter, paging, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
