package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

type CreateFoodStorage interface {
	CreateFood(ctx context.Context, data *model.Food) error
}

type createFoodBiz struct {
	store CreateFoodStorage
}

func NewCreateFoodBiz(store CreateFoodStorage) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateNewFood(ctx context.Context, data *model.Food) error {
	if err := biz.store.CreateFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
