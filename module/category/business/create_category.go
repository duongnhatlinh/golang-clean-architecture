package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

type CreateCategoryStorage interface {
	CreateCategory(ctx context.Context, data *model.Category) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateNewCategory(ctx context.Context, data *model.Category) error {
	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
