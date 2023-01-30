package restaurantlikemodel

import (
	"thucidol/common"
	"time"
)

type User struct {
	common.SimpleUser `json:",inline"`
	LikeAt            *time.Time `json:"created_at,omitempty" gorm:"created_at"`
}
