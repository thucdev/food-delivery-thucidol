package middleware

import (
	"errors"
	"fmt"
	"strings"
	"thucidol/common"
	"thucidol/component/appctx"
	"thucidol/component/tokenprovider/jwt"
	userstore "thucidol/module/user/store"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(err, fmt.Sprintf("wrong authen header"), fmt.Sprintf("ErrWrongAuthHeader"))
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, "")

	if len(parts) < 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(errors.New("wrong header"))
	}

	return parts[1], nil
}

// 1 get token from header
// 2 validate token and parse to payload
// 3 from the token payload, we use user_id to find from DB
func RequireAuth(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenprovider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDbConnection()
		store := userstore.NewSQLStore(db)

		payload, err := tokenprovider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("User has been deleted or banned")))
		}

		user.Mask(false)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
