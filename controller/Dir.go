package controller

import (
	"os"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

// 创建单个文件夹
func Mkdir(c *gin.Context) {
	//获取公共存储的路径
	public_dir := config.Public_path()
	//获取文件夹名称
	name := c.PostForm("name")
	//获取上级路径
	path := c.PostForm("path")

	//判断路径是否带有/
	end_path_str := string(path[len(path)-1])
	if end_path_str != "/" {
		path = path + "/"
	}

	//判断上级路径是否存在
	full_path := public_dir + path

	//需要验证用户传递的path参数是否合法
	if !V_fpath(path) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "参数不合法！",
			"data": "",
		})
		c.Abort()
		return
	}

	//如果路径不存在
	if !V_dir(full_path) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "父级路径不存在！",
			"data": "",
		})
		c.Abort()
		return
	}

	//判断文件夹名称是否合法
	if !V_fname(name) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件夹名称不合法！",
			"data": "",
		})
		c.Abort()
		return
	}

	//准备创建文件夹
	full_dir := public_dir + path + name

	err := os.Mkdir(full_dir, 0755)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "创建失败！",
			"data": err,
		})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "",
		})
	}
}
