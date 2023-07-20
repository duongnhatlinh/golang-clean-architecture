package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	"food_delivery/module/food/storage"
	"food_delivery/pubsub"
)

type HasFood interface {
	GetFoodId() int
}

func RunIncreaseLikeCountAfterUserLikeFood(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes food",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
			likeData := message.Data().(HasFood)
			return store.IncreaseLikeCountFood(ctx, likeData.GetFoodId())
		},
	}
}
