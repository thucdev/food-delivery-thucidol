package main

import (
	"net/http"
	"strconv"
	"thucidol/component/appctx"
	"thucidol/middleware"
	restaurantmodel "thucidol/module/restaurant/model"
	ginrestaurant "thucidol/module/restaurant/transport/gintransport"
	ginuser "thucidol/module/user/transport/ginuser"
	"thucidol/upload/transport/ginupload"

	"github.com/gin-gonic/gin"
)

func setupRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.UploadImage(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/authenticate", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.Profile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data restaurantmodel.Restaurant
		appCtx.GetMainDbConnection().Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		appCtx.GetMainDbConnection().Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
}
