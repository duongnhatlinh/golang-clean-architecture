package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	"food_delivery/module/food/storage"
	"food_delivery/pubsub"
)

func RunDecreaseLikeCountAfterUserDislikeFood(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes food",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
			likeData := message.Data().(HasFood)
			return store.DescendLikeCountFood(ctx, likeData.GetFoodId())
		},
	}
}
