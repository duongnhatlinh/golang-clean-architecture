package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/model"
	"food_delivery/pubsub"
)

type UserDislikeRestaurantStorage interface {
	Delete(ctx context.Context, userId int, restaurantId int) error
}

type userDislikeRestaurantBiz struct {
	store  UserDislikeRestaurantStorage
	pubsub pubsub.Pubsub
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStorage, pubsub pubsub.Pubsub) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, pubsub: pubsub}
}

func (biz *userDislikeRestaurantBiz) UserDislikeRestaurant(ctx context.Context, userId int, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return model.ErrCannotDisLikeRestaurant(err)
	}

	_ = biz.pubsub.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(&model.Restaurant_like{
		RestaurantId: restaurantId,
		UserId:       userId,
	}))

	return nil
}
