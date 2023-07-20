package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurant/storage"
	"food_delivery/pubsub"
)

type HasRestaurant interface {
	GetRestaurantId() int
}

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//	store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
//
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			data := msg.Data().(HasRestaurant)
//
//			_ = store.IncreaseLikeCountRestaurant(ctx, data.GetRestaurantId())
//		}
//	}()
//}

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
			likeData := message.Data().(HasRestaurant)
			return store.IncreaseLikeCountRestaurant(ctx, likeData.GetRestaurantId())
		},
	}
}
