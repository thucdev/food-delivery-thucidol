package restaurantlikemodel

import (
	"fmt"
	"thucidol/common"
	"time"
)

const EntityName = "UsersLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func ErrCanNotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Can not like this restaurant"),
		fmt.Sprintf("ErrCanNotLikeRestaurant"),
	)
}

func ErrCanNotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Can not unlike this restaurant"),
		fmt.Sprintf("ErrCanNotUnlikeRestaurant"),
	)
}
