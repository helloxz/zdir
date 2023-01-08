package router

import (
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
		//获取CID
		cid := c.Request.Header.Get("X-Cid")
		//获取Token
		token := c.Request.Header.Get("X-Token")
		//获取内存中的token
		get_token := string(controller.GetCache([]byte(cid)))

		//如果cid或者token任意一个为空，则终止
		if cid == "" || token == "" {
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "权限不足，请先登录！",
				"data": "",
			})
			c.Abort()
			return
		} else if token == get_token {
			//继续穿透
			c.Next()
		} else {
			//其它任何情况，终止执行
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
	// r.StaticFile("/", "templates/default/index.html")
	// r.StaticFile("/index.html", "templates/default/index.html")
	r.Static("/assets", "templates/assets")
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
	r.GET("/api/user/logout", check_auth(), controller.Logout)

	//上传文件
	r.POST("/api/upload", check_auth(), controller.Upload)
	//删除文件
	r.POST("/api/file/delete", check_auth(), controller.Delete_File)
	//重命名文件
	r.POST("/api/file/rename", check_auth(), controller.RenameFile)

	//创建文件夹
	r.POST("/api/dir/create", check_auth(), controller.Mkdir)
	//跳转到目标URL
	r.GET("/api/jump", controller.JumpURL)

	//后台私有API
	r.POST("/api/option/set", check_auth(), controller.OptionSet)
	r.GET("/api/option/get", check_auth(), controller.OptionGet)
	r.GET("/api/get/app_info", check_auth(), controller.AA_get_app_info)
	r.GET("/api/get/server_info", check_auth(), controller.AA_get_server_infos)
	r.GET("/api/update_sql", check_auth(), controller.UpdateSQL)
	//修改密码
	r.POST("/api/user/change_password", check_auth(), controller.ChangePassword)

	//r.GET("/api/test", controller.Test123)
	//r.GET("/result", controller.OptionGet)

	//前台首页
	r.LoadHTMLFiles("templates/default/index.html", "templates/default/admin.html")
	r.GET("/", controller.DefaultHome)
	r.GET("/index.html", controller.DefaultHome)
	//后台主页
	r.GET("/admin", controller.AdminPage)

	//私有文件路由测试
	private := r.Group("/private", check_auth())
	private.StaticFS("/", http.Dir(public_dir))

	//获取服务端配置
	port := config.Listen()
	r.Run(port)
}
