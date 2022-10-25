package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"zdir/config"
	"zdir/controller"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// 声明一个全局变量用来获取.ini配置
var cfg *ini.File

/*跨域中间件*/
func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token,X-Token,X-Cid")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, HEAD")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

// 验证用户是否登录中间件
func check_auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取Header
		headers := c.Request.Header
		if len(headers["X-Cid"]) >= 1 && len(headers["X-Token"]) >= 1 {
			cid := []byte(headers["X-Cid"][0])
			token := string(headers["X-Token"][0])
			//获取内存
			get_token := string(controller.GetCache(cid))
			fmt.Println(headers)
			fmt.Println("这是header获取的：" + token)
			fmt.Println("这是内存获取的：" + get_token)

			if token == get_token {
				c.Next()
			} else {
				c.JSON(200, gin.H{
					"code": 403,
					"msg":  "权限不足，请先登录！",
					"data": "",
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "权限不足，请先登录！",
				"data": "",
			})
			c.Abort()
			return
		}

	}
}

func Start() {
	//gin运行模式
	RunMode := config.RunMode()
	gin.SetMode(RunMode)

	//日志记录到文件
	f, _ := os.Create("logs/zdir.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//日志同时输出到控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//运行gin
	r := gin.Default()

	//使用跨域中间件
	r.Use(cors())

	public_dir := config.Public_path()

	//为 multipart forms 设置较低的内存限制,8M
	r.MaxMultipartMemory = 8 << 20

	r.GET("/api/filelist", controller.FileList)
	r.StaticFile("/", "data/dist/index.html")
	r.StaticFile("/index.html", "data/dist/index.html")
	r.Static("/assets", "data/dist/assets")
	r.GET("/api/find", controller.Find)
	r.GET("/api/get/appinfo", controller.GetAppInfo)
	r.POST("/api/get/fileinfo", controller.FileInfo)
	r.StaticFS("/public", http.Dir(public_dir))
	//用户状态
	r.GET("/api/user/status", controller.UserStatus)
	//初始化用户
	r.POST("/api/user/init", controller.UserInit)
	//用户登录
	r.POST("/api/user/login", controller.UserLogin)
	r.GET("/api/user/is_login", check_auth(), controller.Is_Login)
	r.GET("/api/user/logout", controller.Logout)

	//上传文件
	r.POST("/api/upload", check_auth(), controller.Upload)
	//删除文件
	r.POST("/api/file/delete", check_auth(), controller.Delete_File)
	//重命名文件
	r.POST("/api/file/rename", check_auth(), controller.RenameFile)

	//创建文件夹
	r.POST("/api/dir/create", check_auth(), controller.Mkdir)

	//获取服务端配置
	port := config.Listen()
	r.Run(port)
}
