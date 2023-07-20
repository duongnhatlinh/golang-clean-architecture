package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

type GetCategoryStorage interface {
	FindCategory(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.Category, error)
}

type getCategoryBiz struct {
	store GetCategoryStorage
}

func NewGetCategoryBiz(store GetCategoryStorage) *getCategoryBiz {
	return &getCategoryBiz{store: store}
}

func (biz *getCategoryBiz) GetCategory(ctx context.Context, id int) (*model.Category, error) {
	data, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(model.EntityName, err)
	}

	return data, nil
}
