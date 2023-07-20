package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/model"
)

type ListUserLikeRestaurantStorage interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		filter *model.Filter,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStorage
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStorage) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}

func (biz *listUserLikeRestaurantBiz) ListUserLikeRestaurant(
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
