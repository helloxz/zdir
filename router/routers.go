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
		context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
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

	r.GET("/api/filelist", controller.FileList)
	r.StaticFile("/", "data/dist/index.html")
	r.StaticFile("/index.html", "data/dist/index.html")
	r.Static("/assets", "data/dist/assets")
	r.GET("/api/find", controller.Find)
	r.GET("/api/get/appinfo", controller.GetAppInfo)
	r.POST("/api/get/fileinfo", controller.FileInfo)
	r.StaticFS("/public", http.Dir(public_dir))

	//获取服务端配置
	port := config.Listen()
	r.Run(port)
}
