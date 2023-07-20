package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/model"
)

type AddFoodStorage interface {
	Create(ctx context.Context, data *model.Cart) error
}

type addFoodBiz struct {
	store AddFoodStorage
}

func NewAddFoodBiz(store AddFoodStorage) *addFoodBiz {
	return &addFoodBiz{store: store}
}

func (biz *addFoodBiz) AddFoodToCart(ctx context.Context, data *model.Cart) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
