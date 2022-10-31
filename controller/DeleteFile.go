package controller

import (
	"fmt"
	"os"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

// 删除文件
func Delete_File(c *gin.Context) {
	//获取文件路径
	fpath := c.PostForm("fpath")
	//取得文件完整路径
	//获取公共存储的路径
	public_dir := config.Public_path()
	full_path := public_dir + fpath
	//如果上传路径为空
	if fpath == "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件路径不能为空！",
			"data": "",
		})
		c.Abort()
		return
	}

	//判断文件路径是否合法
	if !V_fpath(fpath) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件路径不合法！",
			"data": "",
		})
		c.Abort()
		return
	}
	//判断文件或文件夹是否存在
	if !V_is_path(full_path) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件路径不存在！",
			"data": "",
		})
		c.Abort()
		return
	}

	//上述验证通过，执行文件删除

	err := os.RemoveAll(full_path)

	//如果删除失败
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件删除失败！",
			"data": "",
		})
		//打印错误日志
		fmt.Println(err)
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
