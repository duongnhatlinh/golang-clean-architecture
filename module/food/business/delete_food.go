package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

type DeleteFoodStorage interface {
	DeleteFood(
		ctx context.Context,
		condition map[string]interface{},
	) error
	FindFood(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Food, error)
}

type deleteFoodBiz struct {
	store DeleteFoodStorage
}

func NewDeleteFoodBiz(store DeleteFoodStorage) *deleteFoodBiz {
	return &deleteFoodBiz{store: store}
}

func (biz *deleteFoodBiz) RemoveFood(
	ctx context.Context,
	id int,
) error {
	dataOld, err := biz.store.FindFood(ctx, map[string]interface{}{"id": id})

	if dataOld.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.DeleteFood(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
