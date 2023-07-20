package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/model"
	"food_delivery/pubsub"
)

type UserLikeRestaurantStorage interface {
	Create(ctx context.Context, data *model.Restaurant_like) error
}

type userLikeRestaurantBiz struct {
	store  UserLikeRestaurantStorage
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStorage,
	pubsub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:  store,
		pubsub: pubsub,
	}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(ctx context.Context, data *model.Restaurant_like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return model.ErrCannotLikeRestaurant(err)
	}

	_ = biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	return nil
}
