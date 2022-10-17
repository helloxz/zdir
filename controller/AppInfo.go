package controller

import (
	"zdir/config"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体，用来存放文件或文件夹信息
type storages struct {
	Public_domain string
}

// 定义一个结构体，用来存放站点信息
type sites struct {
	Title string
	Name  string
}

// 定义一个结构体，用来返回综合信息
type appinfos struct {
	Storage storages
	Site    sites
}

func GetAppInfo(c *gin.Context) {
	var storage storages

	// 载入配置文件，通过cfg调用
	cfg = config.Load_ini()
	//从配置文件获取存储域名
	public_domain := config.Public_domain()

	//获取请求host
	host := c.Request.Host

	//如果存储域名为空，则使用请求host作为存储域名
	if public_domain == "" {
		storage.Public_domain = "http://" + host + "/public"
	} else {
		storage.Public_domain = public_domain
	}

	//声明结构体
	var site sites
	//获取站点信息
	site.Title, site.Name = config.Site_info()

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
