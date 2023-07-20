package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/model"
)

type GetListCategoriesStorage interface {
	ListCategories(
		ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
	) ([]model.Category, error)
}

type getListCategoriesBiz struct {
	store GetListCategoriesStorage
}

func NewGetListCategoriesBiz(store GetListCategoriesStorage) *getListCategoriesBiz {
	return &getListCategoriesBiz{store: store}
}

func (biz *getListCategoriesBiz) GetListCategories(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
) ([]model.Category, error) {
	data, err := biz.store.ListCategories(ctx, condition, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
