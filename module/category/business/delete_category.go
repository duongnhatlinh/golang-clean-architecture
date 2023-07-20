package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

type DeleteCategoryStorage interface {
	DeleteCategory(
		ctx context.Context,
		condition map[string]interface{},
	) error
	FindCategory(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.Category, error)
}

type deleteCategoryBiz struct {
	store DeleteCategoryStorage
}

func NewDeleteCategoryBiz(store DeleteCategoryStorage) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) RemoveCategory(
	ctx context.Context,
	id int,
) error {
	dataOld, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})

	if dataOld.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.DeleteCategory(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
