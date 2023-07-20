package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
)

type GetFoodStorage interface {
	FindFood(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Food, error)
}

type getFoodBiz struct {
	store GetFoodStorage
}

func NewGetFoodBiz(store GetFoodStorage) *getFoodBiz {
	return &getFoodBiz{store: store}
}

func (biz *getFoodBiz) GetFood(ctx context.Context, id int, moreKeys ...string) (*model.Food, error) {
	data, err := biz.store.FindFood(ctx, map[string]interface{}{"id": id}, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(model.EntityName, err)
	}

	return data, nil
}
