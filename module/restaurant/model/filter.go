package restaurantmodel

type Filter struct {
	UserId int   `json:"user_id" form:user_id`
	Status []int `json:"-"`
}
