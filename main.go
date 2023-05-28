package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/controllers"
	"goweb/global"
	"goweb/initialize"
	"goweb/middlewares"
)

func main() {
	initialize.InitConfig()

	r := gin.Default()
	r.Use(middlewares.Cors())
	login := r.Group("/auth")
	login.POST("/login", controllers.Login)
	login.GET("/test", controllers.Test)

	r.Run(fmt.Sprintf(":%d", global.Settings.Port))
}
