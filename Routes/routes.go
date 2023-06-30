package main

import (
	"gin/App/Controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//建立路由连接
	routes := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8848"} // 设置允许的来源
	routes.Use(cors.New(config))

	//登录功能实现
	loginRoutes := routes.Group("login")
	{
		loginRoutes.POST("/login", Controller.Login)
		loginRoutes.POST("/log", Controller.Log)
	}

	//资源获取
	game := routes.Group("game")
	{
		game.GET("/list", Controller.GameList)
		game.POST("/info", Controller.GameInfo)
	}

	routes.Run()
}
