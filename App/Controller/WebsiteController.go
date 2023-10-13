package Controller

import (
	"gin/App/Models/website"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// WebsiteList 网站列表
func WebsiteList(c *gin.Context)  {
	websiteName := c.Request.FormValue("websiteName")
	page := c.Request.FormValue("page")
	limit := c.Request.FormValue("limit")
	pages, _ := strconv.Atoi(page)
	limits, _ := strconv.Atoi(limit)
	res := website.Website{
		WebsiteName: websiteName,
	}
	list, _ := res.WebsiteList(pages, limits)
	extra := make(map[string]interface{})
	if list != nil {
		extra["count"] = list[0].Count
	} else {
		extra["count"] = 1
	}
	extra["page"] = pages
	c.JSON(http.StatusOK, gin.H{"error": 200, "msg": "", "data": list, "extra": extra})
}