package Controller

import (
	"crypto/md5"
	"gin/App/Models/user"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// Login 账号注册
func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	image := c.Request.FormValue("image")
	account := c.Request.FormValue("account")
	pwd := c.Request.FormValue("password")
	h := md5.New()
	password, err := h.Write([]byte(pwd))
	if err != nil {
		log.Fatal(err)
	}
	model := user.User{
		Name:       name,
		Image:      image,
		Account:    account,
		Password:   strconv.Itoa(password),
		CreateTime: int(time.Now().Unix()),
		UpdateTime: int(time.Now().Unix()),
	}
	ids := model.Login()
	c.JSON(200, gin.H{"code": 200, "msg": ids, "data": ""})
}

// Log 账号登录
func Log(c *gin.Context) {
	account := c.Request.FormValue("account")
	pwd := c.Request.FormValue("password")
	h := md5.New()
	password, err := h.Write([]byte(pwd))
	if err != nil {
		log.Fatal(err)
	}
	model := user.User{
		Account:  account,
		Password: strconv.Itoa(password),
	}
	info := model.Log()
	c.JSON(200, gin.H{"code": 200, "msg": info, "data": ""})
}
