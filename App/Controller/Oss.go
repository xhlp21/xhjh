package Controller

import (
	"gin/App/Service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func File(c *gin.Context) {
	files, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err.Error())
	}
	code, url := Service.UploadToQiNiu(files)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "OK",
		"url":  url,
	})
}
