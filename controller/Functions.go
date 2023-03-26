// 常用方法集合
package controller

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// md5字符串加密
func md5s(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成一个随机字符串
func RandStr(n int) string {
	//如果seed固定，那么每次程序重启后重新生成随机数会重复上一次的随机数，所以这里设置一个随机seed
	rand.Seed(time.Now().UnixNano())
	var bytes []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Int31()%62]
	}
	return string(result)
}

// 密码加密,传递用户名 + 密码，最后返回加密后的密码
func password_encry(username string, password string) string {
	password_encry := md5s(username + password)
	return password_encry
}

// 获取客户端IP
func GetClientIp(c *gin.Context) string {
	//尝试通过X-Forward-For获取
	ip := c.Request.Header.Get("X-Forward-For")
	//如果没获取到，则通过X-real-ip获取
	if ip == "" {
		ip = c.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		//依然没获取到，则通过gin自身方法获取
		ip = c.ClientIP()
	}
	//判断IP格式是否正确，避免伪造IP
	if V_ip(ip) {
		return ip
	} else {
		return "0.0.0.0"
	}
}

// url编码处理，有问题，暂时不要用
func UrlEncode(content string) string {
	//默认编码处理
	new_url := url.QueryEscape(content)
	//字符串替换
	new_url = strings.Replace(new_url, "!", "%21", -1)
	new_url = strings.Replace(new_url, "'", "%27", -1)
	new_url = strings.Replace(new_url, "(", "%28", -1)
	new_url = strings.Replace(new_url, ")", "%29", -1)
	new_url = strings.Replace(new_url, "*", "%2A", -1)
	//new_url = strings.Replace(new_url, "\\s", "%20", -1)
	new_url = strings.Replace(new_url, "%2F", "/", -1)
	new_url = strings.Replace(new_url, "+", "%2B", -1)
	new_url = strings.Replace(new_url, " ", "%20", -1)

	return new_url
}
