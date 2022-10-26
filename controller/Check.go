// 检查各种信息，比如用户信息
package controller

import (
	"zdir/config"

	"github.com/gin-gonic/gin"
)

func UserStatus(c *gin.Context) {
	// 获取用户、密码
	username, password := config.User_info()
	//如果用户名、密码不为空，则返回真
	if username != "" && password != "" {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": "",
		})
		c.Abort()
		return
	} else {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  "Not ready",
			"data": "",
		})
	}
}
