package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurant/storage"
	"food_delivery/pubsub"
)

//	func DescendLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//		c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserDislikeRestaurant)
//		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
//
//		go func() {
//			defer common.AppRecover()
//			for {
//				msg := <-c
//				data := msg.Data().(int)
//
//				_ = store.DescendLikeCountRestaurant(ctx, data)
//			}
//		}()
//	}
func RunDecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
			likeData := message.Data().(HasRestaurant)
			return store.DescendLikeCountRestaurant(ctx, likeData.GetRestaurantId())
		},
	}
}
