package controller

import (
	"encoding/base64"
	"fmt"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

// 该函数接收base64的字符串，然后进行接码并302重定向
func JumpURL(c *gin.Context) {
	urlstr := c.Query("urlstr")
	// 解码
	s, err := base64.StdEncoding.DecodeString(urlstr)

	if err != nil {
		fmt.Println("base64 decode error:", err)
		return
	}
	public_domain := config.Public_domain(c)
	source_path := string(s)
	url := public_domain + source_path
	//c.Header("Content-Disposition", "attachment; filename=\""+source_path+"\"")
	c.Redirect(302, url)
	c.Abort()
}
