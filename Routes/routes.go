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
		game.GET("/list/recommend", Controller.GameListRecommend)
		game.POST("/info", Controller.GameInfo)
		game.GET("/list", Controller.GameList)
	}

	//游戏类型
	types := routes.Group("type")
	{
		types.GET("/list",Controller.TypeList)
		types.GET("/game",Controller.TypeGame)

	}

	//网站网址
	website := routes.Group("website")
	{
		website.GET("/list",Controller.WebsiteList)
	}

	//云文件上传
	oss := routes.Group("oss")
	{
		oss.POST("/files",Controller.File)
	}

	routes.Run()
}
