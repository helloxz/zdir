package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"zdir/cli"
	"zdir/model"

	"github.com/gin-gonic/gin"
)

// 默认首页
func DefaultHome(c *gin.Context) {
	//获取站点信息
	site_data := model.OptionGet("site_data")

	//interface{}类型转为string类型
	custom_header := fmt.Sprintf("%v", site_data["custom_header"])

	//网站标题
	title := site_data["title"]
	//如果没获取到网站标题，则默认使用Zdir作为标题
	if title == nil {
		title = "Zdir"
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       title,
		"keywords":    site_data["keywords"],
		"description": site_data["description"],
		//template.HTML()防止HTML被转义
		"custom_header": template.HTML(custom_header),
		"version":       cli.VersionDate,
	})
}

// 后台主页
func AdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"version": cli.VersionDate,
	})
}
