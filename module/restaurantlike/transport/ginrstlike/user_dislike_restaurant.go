package ginrstlike

import (
	"net/http"
	"thucidol/common"
	"thucidol/component/appctx"
	restaurantlikebiz "thucidol/module/restaurantlike/biz"
	restaurantlikestore "thucidol/module/restaurantlike/store"

	"github.com/gin-gonic/gin"
)

func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInValidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestore.NewSQLStore(appCtx.GetMainDbConnection())
		// decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())

		biz := restaurantlikebiz.NewUserDislikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
