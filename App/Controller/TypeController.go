package Controller

import (
	"gin/App/Models/type"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TypeList 类型列表
func TypeList(c *gin.Context)  {
	res := _type.Type{}
	list := res.TypeList()
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "", "data": list})
}

// TypeGame 游戏类型列表
func TypeGame(c *gin.Context)  {
	typeId := c.Request.FormValue("type_id")
	if typeId == ""{
		c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "", "data": ""})
		return
	}
	res := _type.Type{}
	list, err := res.TypeGame(typeId)
	if err != nil{
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "", "data": list})
}