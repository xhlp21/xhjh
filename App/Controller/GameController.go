package Controller

import (
	"gin/App/Models/game"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GameListRecommend 查出筛选的游戏推荐
func GameListRecommend(c *gin.Context) {
	res := game.Game{}
	list := res.GameListRecommend()
	c.JSON(http.StatusOK, gin.H{"error": 200, "msg": "", "data": list})
}

// GameInfo 进入到游戏的详情页
func GameInfo(c *gin.Context) {
	gameId := c.Request.FormValue("gameId")
	ids, err := strconv.Atoi(gameId)
	if err != nil {
		log.Fatal(err.Error())
	}
	if gameId == "" {
		c.JSON(http.StatusOK, gin.H{"error": 500, "msg": "传递ID不能为空！", "data": ""})
	}
	res := game.Game{
		GameId: ids,
	}
	info := res.GameInfo()
	c.JSON(http.StatusOK, gin.H{"error": 200, "msg": "", "data": info})
}

// GameList 全部游戏列表
func GameList(c *gin.Context) {
	gameName := c.Request.FormValue("gameName")
	res := game.Game{
		GameName: gameName,
	}
	list, _ := res.GameList()
	c.JSON(http.StatusOK, gin.H{"error": 200, "msg": "", "data": list})
}
