package main

import (
	"thucidol/component/appctx"
	"thucidol/middleware"
	ginuser "thucidol/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func setupAdminRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin", middleware.RequireAuth(appCtx), middleware.RoleRequired(appCtx, "admin"))

	{
		admin.GET("/profile", ginuser.Profile(appCtx))
	}
}
