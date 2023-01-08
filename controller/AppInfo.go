package controller

import (
	"fmt"
	"zdir/config"
	"zdir/model"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体，用来存放文件或文件夹信息
type storages struct {
	Public_domain string
}

// 定义一个结构体，用来存放站点信息
type sites struct {
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
}

// 定义一个结构体，用来返回综合信息
type appinfos struct {
	Storage storages
	Site    sites
}

func GetAppInfo(c *gin.Context) {
	var storage storages

	//从配置文件获取存储域名
	public_domain := config.Public_domain(c)

	storage.Public_domain = public_domain

	//声明结构体
	var site sites
	//获取站点信息
	site_data := model.OptionGet("site_data")
	if site_data["title"] == nil {
		site.Title = "Zdir"
	} else {
		site.Title = fmt.Sprintf("%v", site_data["title"])
	}

	site.Logo = fmt.Sprintf("%v", site_data["logo"])
	site.Keywords = fmt.Sprintf("%v", site_data["keywords"])
	site.Description = fmt.Sprintf("%v", site_data["description"])

	//声明一个map
	var appinfo appinfos
	appinfo.Storage = storage
	appinfo.Site = site

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": appinfo,
	})
}
