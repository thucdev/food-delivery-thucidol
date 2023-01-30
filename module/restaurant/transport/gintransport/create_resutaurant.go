package ginrestaurant

import (
	"net/http"

	"thucidol/common"
	"thucidol/component/appctx"
	restaurantbiz "thucidol/module/restaurant/biz"
	restaurantmodel "thucidol/module/restaurant/model"
	restaurantstorage "thucidol/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDbConnection()
		var data restaurantmodel.RestaurantCreate
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		//defer to get error when project crash
		// go func ()  {
		// 	defer common.AppRecover()

		// 	arr := []int{}
		// 	log.Println(arr[0])
		// }()

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
