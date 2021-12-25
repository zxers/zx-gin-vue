package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zxers/zx-gin-vue/middleware"
	"github.com/zxers/zx-gin-vue/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//r.Static("/static", "static")
	//r.LoadHTMLGlob("template/*")

	api := r.Group("v1")
	{
		api.POST("/user/register", controller.Register)
		api.POST("/user/login", controller.Login)
		api.GET("/user/info", middleware.AuthMiddleware(),controller.Info)
	}
	
	return r
}