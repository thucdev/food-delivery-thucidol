package restaurantmodel

import (
	"errors"
	"strings"
	"thucidol/common"
	usermodel "thucidol/module/user/model"
)

type RestaurantType string

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false"`
	Type            RestaurantType     `json:"type" gorm:"column:type;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	LikedCount      int                `json:"liked_count" gorm:"liked_count"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdmin)
	}
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Logo  string         `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string          `json:"name" gorm:"column:name;"`
	Addr            string          `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image   `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images  `json:"cover" gorm:"column:cover;"`
	User            *usermodel.User `json:"user" gorm:"preload:false"`
	UserId          int             `json:"-" gorm:"column:user_id;"`
}

func (data *RestaurantCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
