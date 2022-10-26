package controller

import (
	"os"
	"regexp"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	//获取文件路径
	path := c.PostForm("path")
	//如果上传路径为空
	if path == "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "上传路径不能为空！",
			"data": "",
		})
		c.Abort()
		return
	}
	//如果上传路径不合法
	//判断用户传递的路径是否合法
	var validPath = regexp.MustCompile(`^(\.|\..).+`)
	v_re := validPath.MatchString(path)
	if v_re {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件夹名称不合法！",
			"data": "",
		})
		c.Abort()
		return
	}

	//判断上传路径是否带有/
	end_path_str := string(path[len(path)-1])
	if end_path_str != "/" {
		path = path + "/"
	}

	//组合公共文件完整路径
	public_dir := config.Public_path()
	full_path := public_dir + path

	//判断文件是否存在，如果不存在，则终止执行
	_, err := os.Stat(full_path)
	if os.IsNotExist(err) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件夹不存在！",
			"data": "",
		})
		c.Abort()
		return
	}

	// 单文件
	file, _ := c.FormFile("file")

	dst := full_path + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)
	//返回上传成功
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "",
	})
}
