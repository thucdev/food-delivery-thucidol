package ginrstlike

import (
	"net/http"
	"thucidol/common"
	"thucidol/component/appctx"
	restaurantlikebiz "thucidol/module/restaurantlike/biz"
	restaurantlikemodel "thucidol/module/restaurantlike/model"
	restaurantlikestore "thucidol/module/restaurantlike/store"

	"github.com/gin-gonic/gin"
)

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInValidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInValidRequest(err))
		}

		paging.Fullfil()

		store := restaurantlikestore.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}

}
