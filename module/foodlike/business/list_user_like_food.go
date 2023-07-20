package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/foodlike/model"
)

type ListUserLikeFoodStorage interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		filter *model.Filter,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeFoodBiz struct {
	store ListUserLikeFoodStorage
}

func NewListUserLikeFoodBiz(store ListUserLikeFoodStorage) *listUserLikeFoodBiz {
	return &listUserLikeFoodBiz{store: store}
}

func (biz *listUserLikeFoodBiz) ListUserLikeFood(
	ctx context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	filter *model.Filter,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	data, err := biz.store.List(ctx, conditions, paging, filter, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
