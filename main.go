package main

import (
	"goweb/controllers"
	"goweb/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	login := r.Group("/auth")
	login.POST("/login", controllers.Login)
	login.GET("/test", controllers.Test)
	r.Run()
}
