package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/model"
)

type UpdateInfoUserStorage interface {
	UpdateUser(ctx context.Context, id int, data *model.UserUpdate) error
}

type updateInfoUserBiz struct {
	store UpdateInfoUserStorage
}

func NewUpdateInfoUserBiz(store UpdateInfoUserStorage) *updateInfoUserBiz {
	return &updateInfoUserBiz{store: store}
}

func (biz *updateInfoUserBiz) UpdateInfoUser(ctx context.Context, id int, data *model.UserUpdate) error {
	err := biz.store.UpdateUser(ctx, id, data)

	if err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
