package main

import (
	"gin/App/Controller"
	"github.com/gin-gonic/gin"
)

func main() {

	//建立路由连接
	routes := gin.Default()

	loginRoutes := routes.Group("login")
	{
		loginRoutes.POST("/login", Controller.Login)
		loginRoutes.POST("/log", Controller.Log)
	}

	routes.Run()
}
