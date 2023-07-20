package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

type UpdateCategoryStorage interface {
	UpdateCategory(
		ctx context.Context,
		condition map[string]interface{},
		data *model.UpdateCategory,
	) error
	FindCategory(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.Category, error)
}

type updateCategoryBiz struct {
	store UpdateCategoryStorage
}

func NewUpdateCategoryBiz(store UpdateCategoryStorage) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) EditCategory(
	ctx context.Context,
	id int,
	data *model.UpdateCategory,
) error {
	dataOld, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})

	if dataOld.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.UpdateCategory(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
