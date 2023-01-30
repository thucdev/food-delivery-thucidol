package middleware

import (
	"errors"
	"thucidol/common"
	"thucidol/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		hashFound := false
		u := c.MustGet(common.CurrentUser).(common.Requester)
		for _, role := range allowRoles {
			if u.GetRole() == role {
				hashFound = true
			}
		}
		if hashFound {
			c.Next()
		} else {
			panic(common.ErrNoPermission(errors.New("invalid role")))
		}

	}
}
