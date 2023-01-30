package ginuser

import (
	"net/http"
	"thucidol/common"
	"thucidol/component/appctx"
	"thucidol/component/hasher"
	"thucidol/component/tokenprovider/jwt"
	userbiz "thucidol/module/user/biz"
	usermodel "thucidol/module/user/model"
	userstore "thucidol/module/user/store"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInValidRequest(err))
		}

		db := appCtx.GetMainDbConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
