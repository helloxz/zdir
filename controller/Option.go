package controller

import (
	"fmt"
	"zdir/model"

	"github.com/gin-gonic/gin"
)

func OptionSet(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	note := c.PostForm("note")

	//设置数据
	RowsAffected, err := model.OptionSet(key, value, note)
	if err != nil {
		//设置失败
		fmt.Println(err)

		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "Failed to set",
			"data": "",
		})
		return
	}

	//设置成功，返回数据
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": RowsAffected,
	})
}

// 根据key查询选项参数
func OptionGet(c *gin.Context) {
	//获取key参数
	key := c.Query("key")
	if key == "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "Missing parameter!",
			"data": "",
		})
		return
	}

	//查询数据库
	data := model.OptionGet(key)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
