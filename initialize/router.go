package initialize

import (
	"goweb/middlewares"
	"goweb/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.GinLogger(), middlewares.Cors(), middlewares.GinRecovery(true))
	//路由分组
	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	return Router
}
