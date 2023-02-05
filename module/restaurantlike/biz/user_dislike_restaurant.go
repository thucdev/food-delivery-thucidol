package restaurantlikebiz

import (
	"context"
	"log"
	"thucidol/common"
	restaurantlikemodel "thucidol/module/restaurantlike/model"
	"thucidol/pubsub"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

// type DecLikedCountResStore interface {
// 	DecreaseLikeCount(ctx context.Context, id int) error
// }

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
	// decStore DecLikedCountResStore
	ps pubsub.PubSub
}

// func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, decStore DecLikedCountResStore) *userDislikeRestaurantBiz {
func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, ps pubsub.PubSub) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, ps: ps}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCanNotUnlikeRestaurant(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserDislikeRestaurant,
		pubsub.NewMessage(&restaurantlikemodel.Like{RestaurantId: restaurantId})); err != nil {
		log.Panicln(err)
	}

	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	// go func() {
	// 	defer common.AppRecover()
	// 	if err := biz.decStore.DecreaseLikeCount(ctx, restaurantId); err != nil {
	// 		panic(err)
	// 	}

	// }()

	return nil
}
