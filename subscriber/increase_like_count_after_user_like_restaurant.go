package subscriber

import (
	"context"
	"log"
	"thucidol/component/appctx"
	restaurantstorage "thucidol/module/restaurant/storage"
	"thucidol/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

// func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}

func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification when user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Panicln("restaurant id:", likeData.GetRestaurantId())
			return nil
		},
	}
}
