package ginrestaurant

import (
	"net/http"

	"thucidol/common"
	"thucidol/component/appctx"
	restaurantbiz "thucidol/module/restaurant/biz"
	restaurantmodel "thucidol/module/restaurant/model"
	restaurantrepo "thucidol/module/restaurant/repository"
	restaurantstorage "thucidol/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInValidRequest(err))
		}

		pagingData.Fullfil()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInValidRequest(err))

		}

		db := appCtx.GetMainDbConnection()

		store := restaurantstorage.NewSQLStore(db)
		// likeStore := restaurantlikestore.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
