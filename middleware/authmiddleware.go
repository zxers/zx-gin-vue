package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zxers/zx-gin-vue/common"
	"github.com/zxers/zx-gin-vue/model"
	"github.com/zxers/zx-gin-vue/db"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": 	"权限不够",
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": 	"权限不够",
			})
			ctx.Abort()
			return
		}
		userId := claims.UserId
		var user model.User
		db.DB.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":	"权限不够",
			})
		}

		//用户存在，将user的信息写入context
		ctx.Set("user", user)
		ctx.Next()
	}
}

