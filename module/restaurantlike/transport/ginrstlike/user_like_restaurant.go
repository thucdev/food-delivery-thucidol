package ginrstlike

import (
	"net/http"
	"thucidol/common"
	"thucidol/component/appctx"
	restaurantstorage "thucidol/module/restaurant/storage"
	restaurantlikebiz "thucidol/module/restaurantlike/biz"
	restaurantlikemodel "thucidol/module/restaurantlike/model"
	restaurantlikestore "thucidol/module/restaurantlike/store"

	"github.com/gin-gonic/gin"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInValidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestore.NewSQLStore(appCtx.GetMainDbConnection())
		incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		// not important, put it in go routin
		go func() {
			defer common.AppRecover()
			if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
				panic(err)
			}

		}()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
