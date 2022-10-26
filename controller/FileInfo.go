package controller

import (
	"log"
	"os"
	"strings"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体，用来存放文件或文件夹信息
type finfo struct {
	Name  string
	Size  int64
	Mtime string
	Fpath string
	Ext   string
}

func FileInfo(c *gin.Context) {
	//获取公共存储的路径
	public_dir := config.Public_path()

	//声明结构体
	var new_info finfo

	//获取路径参数
	fpath := string(c.PostForm("fpath"))

	//判断用户传递的路径是否合法
	v_re := !V_fpath(fpath)
	if v_re {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件夹不合法！",
			"data": "",
		})
		c.Abort()
		return
	}
	//拼接完整路径
	full_path := public_dir + fpath

	//获取文件信息
	finfo, err := os.Stat(full_path)

	//fmt.Println(full_path)

	//如果出现错误，比如文件夹不存在
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err,
			"data": "",
		})
		c.Abort()
		return
	} else {
		//如果是目录
		if finfo.IsDir() {
			c.JSON(200, gin.H{
				"code": -1000,
				"msg":  "只允许文件参数！",
				"data": "",
			})
			c.Abort()
			return
		} else {
			//获取扩展名
			ext_temp := strings.Split(fpath, ".")
			//取分隔的最后一个元素
			new_info.Ext = strings.ToLower(ext_temp[len(ext_temp)-1])
			//获取文件修改时间
			new_info.Mtime = finfo.ModTime().Format("2006-01-02 15:04:05")
			//获取文件大小
			new_info.Size = finfo.Size()
			//获取文件名称
			new_info.Name = finfo.Name()
			//获取路径
			new_info.Fpath = fpath
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "success",
				"data": new_info,
			})
			return
		}

	}
}
