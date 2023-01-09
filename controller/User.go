// 用户行为操作
package controller

import (
	"strings"
	"time"
	"zdir/config"
	"zdir/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
		v_re := V_username(username)
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
		v_re = V_password(password)
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
		password = password_encry(username, password)
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

// 修改用户名、密码
func ChangePassword(c *gin.Context) {
	//获取表单数据
	f_username := c.PostForm("username")
	//去除空白字符
	f_username = strings.Replace(f_username, " ", "", -1)
	//用户名转小写
	f_username = strings.ToLower(f_username)
	f_old_password := c.PostForm("old_password")
	f_new_password := c.PostForm("new_password")
	//去除密码空白字符
	f_new_password = strings.Replace(f_new_password, " ", "", -1)
	f_confirm_password := c.PostForm("confirm_password")

	//验证用户名是否合法
	if !V_username(f_username) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "用户名不符合规范！",
			"data": "",
		})
		return
	}
	//验证密码是否合法
	if !V_password(f_new_password) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "新密码不符合规范！",
			"data": "",
		})
		return
	}
	//判断2次密码是否一致
	if f_new_password != f_confirm_password {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "两次密码不一致！",
			"data": "",
		})
		return
	}

	//从配置文件获取原始密码
	config_password := viper.GetString("users.password")
	//从配置文件获取用户名
	config_username := viper.GetString("users.username")
	//加密用户原密码
	encry_old_password := password_encry(config_username, f_old_password)
	//如果原始密码和用户提供的密码不一致，则终止执行
	if config_password != encry_old_password {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "原始密码不正确！",
			"data": "",
		})
		return
	}

	//加密新的密码
	encry_new_password := password_encry(f_username, f_new_password)
	//写入保存
	is_bool := SetKVS("users.username", f_username)
	SetKVS("users.password", encry_new_password)

	//如果写入失败了
	if !is_bool {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "配置文件写入失败！",
			"data": "",
		})
		return
	} else {
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

	//如果用户输入的密码是空的，则不允许登录
	if get_password == "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "Password cannot be empty!",
			"data": "",
		})
		return
	}

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

		//获取IP、ua等信息写入日志记录
		ip := GetClientIp(c)
		ua := c.Request.Header.Get("User-Agent")
		//当前时间
		now := time.Now()
		data := model.Z_login_log{
			Ip:         ip,
			Ua:         ua,
			Cid:        cid,
			Token:      token,
			Expired_at: now.Add(time.Hour * 24 * 7).Unix(),
			Behavior:   "login",
			State:      1,
		}
		//插入数据库
		model.LoginLogInsert(data)
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
