package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/foodlike/model"
	"food_delivery/pubsub"
)

type UserLikeFoodStorage interface {
	Create(ctx context.Context, data *model.Food_like) error
}

type userLikeFoodBiz struct {
	store  UserLikeFoodStorage
	pubsub pubsub.Pubsub
}

func NewUserLikeFoodBiz(
	store UserLikeFoodStorage,
	pubsub pubsub.Pubsub,
) *userLikeFoodBiz {
	return &userLikeFoodBiz{
		store:  store,
		pubsub: pubsub,
	}
}

func (biz *userLikeFoodBiz) UserLikeFood(ctx context.Context, data *model.Food_like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return model.ErrCannotLikeFood(err)
	}

	_ = biz.pubsub.Publish(ctx, common.TopicUserLikeFood, pubsub.NewMessage(data))

	return nil
}
