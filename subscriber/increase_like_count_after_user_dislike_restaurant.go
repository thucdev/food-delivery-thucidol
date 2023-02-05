package subscriber

import (
	"context"
	"thucidol/component/appctx"
	restaurantstorage "thucidol/module/restaurant/storage"
	"thucidol/pubsub"
)

// func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user dislike restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}
