// 用户行为操作
package controller

import (
	"regexp"
	"strings"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

// 声明一个结构体，用来返回认证信息
type Authorization struct {
	UserName string
	Cid      string
	Token    string
}

// 用户初始化函数
func UserInit(c *gin.Context) {
	// 获取配置文件中的用户名、密码
	username, password := config.User_info()

	if username != "" && password != "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "用户已经初始化，无需再次操作！",
			"data": "",
		})
		c.Abort()
		return
	} else {
		//获取POST Form中的数据
		username = c.PostForm("username")
		password = c.PostForm("password")
		//去除空白字符串
		username = strings.Replace(username, " ", "", -1)
		password = strings.Replace(password, " ", "", -1)

		//判断用户名是否符合规范
		var validUser = regexp.MustCompile(`^[a-z0-9]{2,16}`)
		v_re := validUser.MatchString(username)
		if !v_re {
			c.JSON(200, gin.H{
				"code": -1000,
				"msg":  "用户名不符合规范！",
				"data": "",
			})
			c.Abort()
			return
		}

		//判断密码是否符合规范
		var validPass = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&\*\(\)_\.]{8,16}`)
		v_re = validPass.MatchString(password)
		if !v_re {
			c.JSON(200, gin.H{
				"code": -1000,
				"msg":  "密码不符合规范！",
				"data": "",
			})
			c.Abort()
			return
		}

		//密码进行md5加密
		password = md5s(username + password)
		//用户名转小写
		username = strings.ToLower(username)

		SetKVS("users.username", username)
		SetKVS("users.password", password)

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "",
		})
	}

}

// 用户登录函数
func UserLogin(c *gin.Context) {
	// 获取配置文件中的用户名、密码
	username, password := config.User_info()
	//如果用户名或密码其中一个为空，则需要先初始化
	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "用户名或密码为空，请选初始化！",
			"data": "",
		})
		c.Abort()
		return
	}

	//获取POST Form中的数据
	get_username := c.PostForm("username")
	get_password := c.PostForm("password")
	//去除空白字符串
	username = strings.Replace(username, " ", "", -1)
	password = strings.Replace(password, " ", "", -1)

	//密码进行md5加密
	get_password = md5s(get_username + get_password)
	//用户名转小写
	get_username = strings.ToLower(username)

	//比较用户名、密码是否匹配
	if username == get_username && password == get_password {
		//生成一个6位CID
		cid := RandStr(6)
		//生成一个随机token,用户名 + 加密密码 + 6位随机字符串
		token := md5s(get_username + get_password + RandStr(6))
		//设置cookie，有效期7天，不需要HTTPONLY
		ttl := 60 * 60 * 24 * 7
		// c.SetCookie("USERNAME", get_username, ttl, "/", "", false, false)
		// c.SetCookie("CID", cid, ttl, "/", "", false, false)
		// c.SetCookie("TOKEN", token, ttl, "/", "", false, false)

		var auth Authorization
		auth.Cid = cid
		auth.UserName = get_username
		auth.Token = token

		//token保存到内存中，以便下次直接验证使用
		SetCache([]byte(cid), []byte(token), ttl)

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": auth,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "用户名或密码错误！",
			"data": "",
		})
		c.Abort()
		return
	}
}

// 用户点击退出
func Logout(c *gin.Context) {
	//获取用户CID
	cid := c.Query("cid")
	//判断CID是否合法
	if !V_cid(cid) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "CID不合法！",
			"data": "",
		})
		c.Abort()
		return
	}

	//继续执行
	//删除认证缓存
	DelCache(cid)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "",
	})
}

// 检查用户是否登录
func Is_Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "yes",
		"data": "",
	})
}
