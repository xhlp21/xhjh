package Controller

import (
	"gin/App/Models/game"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GameList 查出筛选的游戏推荐
func GameList(c *gin.Context) {
	res := game.Game{}
	list := res.GameList()
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
