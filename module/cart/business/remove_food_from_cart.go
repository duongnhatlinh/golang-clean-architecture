package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/model"
)

type RemoveFoodStorage interface {
	Delete(ctx context.Context, userId int, foodId int) error
	FindFood(ctx context.Context, userId int, foodId int) (*model.Cart, error)
}

type removeFoodBiz struct {
	store RemoveFoodStorage
}

func NewRemoveFoodBiz(store RemoveFoodStorage) *removeFoodBiz {
	return &removeFoodBiz{store: store}
}

func (biz *removeFoodBiz) RemoveFoodFromCart(ctx context.Context, userId int, foodId int) error {
	f, err := biz.store.FindFood(ctx, userId, foodId)

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if f.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.Delete(ctx, userId, foodId); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
