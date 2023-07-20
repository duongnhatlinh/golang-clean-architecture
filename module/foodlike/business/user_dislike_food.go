package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/foodlike/model"
	"food_delivery/pubsub"
)

type UserDislikeFoodStorage interface {
	Delete(ctx context.Context, userId int, FoodId int) error
}

type userDislikeFoodBiz struct {
	store  UserDislikeFoodStorage
	pubsub pubsub.Pubsub
}

func NewUserDislikeFoodBiz(store UserDislikeFoodStorage, pubsub pubsub.Pubsub) *userDislikeFoodBiz {
	return &userDislikeFoodBiz{store: store, pubsub: pubsub}
}

func (biz *userDislikeFoodBiz) UserDislikeFood(ctx context.Context, userId int, foodId int) error {
	err := biz.store.Delete(ctx, userId, foodId)

	if err != nil {
		return model.ErrCannotDisLikeFood(err)
	}

	_ = biz.pubsub.Publish(ctx, common.TopicUserDislikeFood, pubsub.NewMessage(&model.Food_like{
		FoodId: foodId,
		UserId: userId,
	}))

	return nil
}
