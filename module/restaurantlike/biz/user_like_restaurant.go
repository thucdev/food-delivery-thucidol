package restaurantlikebiz

import (
	"context"
	"log"
	"thucidol/component/asyncjob"
	restaurantlikemodel "thucidol/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountResStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncLikedCountResStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCanNotLikeRestaurant(err)
	}

	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	// if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
	// 	log.Println(err)
	// }

	return nil
}
