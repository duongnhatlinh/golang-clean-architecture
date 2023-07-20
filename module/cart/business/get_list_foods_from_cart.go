package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/model"
)

type GetListFoodsStorage interface {
	List(ctx context.Context, userId int, moreKeys ...string) ([]model.Cart, error)
}

type getListFoodsBiz struct {
	store GetListFoodsStorage
}

func NewGetListFoodsBiz(store GetListFoodsStorage) *getListFoodsBiz {
	return &getListFoodsBiz{store: store}
}

func (biz *getListFoodsBiz) GetListFoodsFromCart(
	ctx context.Context,
	userId int,
	moreKeys ...string,
) ([]model.Cart, error) {
	data, err := biz.store.List(ctx, userId, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
