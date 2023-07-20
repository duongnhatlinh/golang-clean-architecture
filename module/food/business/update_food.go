package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

type UpdateFoodStorage interface {
	UpdateFood(
		ctx context.Context,
		condition map[string]interface{},
		data *model.UpdateFood,
	) error
	FindFood(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Food, error)
}

type updateFoodBiz struct {
	store UpdateFoodStorage
}

func NewUpdateFoodBiz(store UpdateFoodStorage) *updateFoodBiz {
	return &updateFoodBiz{store: store}
}

func (biz *updateFoodBiz) EditFood(
	ctx context.Context,
	id int,
	data *model.UpdateFood,
) error {
	dataOld, err := biz.store.FindFood(ctx, map[string]interface{}{"id": id})

	if dataOld.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.UpdateFood(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
