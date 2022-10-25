package controller

import (
	"fmt"
	"os"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

// 重命名文件
func RenameFile(c *gin.Context) {
	//获取公共存储的路径
	public_dir := config.Public_path()
	//获取原文件名
	old_name := c.PostForm("old_name")
	//获取新文件名
	new_name := c.PostForm("new_name")
	//获取文件路径
	fpath := c.PostForm("fpath")

	//判断路径是否带有/
	end_path_str := string(fpath[len(fpath)-1])
	if end_path_str != "/" {
		fpath = fpath + "/"
	}

	//判断文件名是否一样
	if old_name == new_name {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "新旧文件名一致，无需修改！",
			"data": "",
		})
		c.Abort()
		return
	}
	//新文件名不能为空
	if new_name == "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "新文件名不能为空！",
			"data": "",
		})
		c.Abort()
		return
	}
	//判断路径是否是文件夹,并且验证路径是否合法
	if !V_dir(public_dir+fpath) || !V_fpath(public_dir+fpath) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "不是有效的文件路径！",
			"data": "",
		})
		c.Abort()
		return
	}

	//拼接原文件名完整路径
	old_path := public_dir + fpath + old_name
	new_path := public_dir + fpath + new_name
	//判断原文件是否存在
	if !V_is_file(old_path) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件路径不存在！",
			"data": "",
		})
		c.Abort()
		return
	}

	//判断新的文件名是否合法，文件名不应该包含/或\或|
	if !V_fname(new_name) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件名称不合法！",
			"data": "",
		})
		c.Abort()
		return
	}

	//验证全部通过，进行文件重命名

	err := os.Rename(old_path, new_path)
	// fmt.Println(old_path)
	// fmt.Println(new_path)

	if err != nil {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件重命名失败！",
			"data": err,
		})
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
