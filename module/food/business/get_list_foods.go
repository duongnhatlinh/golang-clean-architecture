package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

type GetListFoodsStorage interface {
	ListFood(
		ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
		filter *model.Filter,
		moreKeys ...string,
	) ([]model.Food, error)
}

type getListFoodsBiz struct {
	store GetListFoodsStorage
}

func NewGetListFoodsBiz(store GetListFoodsStorage) *getListFoodsBiz {
	return &getListFoodsBiz{store: store}
}

func (biz *getListFoodsBiz) GetListFoods(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
	filter *model.Filter,
	moreKeys ...string,
) ([]model.Food, error) {
	data, err := biz.store.ListFood(ctx, condition, paging, filter, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
