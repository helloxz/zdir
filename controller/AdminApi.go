// 管理员API，全部需要验证，后台操作的相关方法都写到这里，统一要求AA_开头
package controller

import (
	"runtime"
	"zdir/cli"
	"zdir/config"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/spf13/viper"
)

// 声明一个结构体，用于格式化配置文件信息
type AppInfo struct {
	Version       string `json:"version"`
	Port          string `json:"port"`
	Runmode       string `json:"runmode"`
	Public_domain string `json:"public_domain"`
	Public_path   string `json:"public_path"`
	Username      string `json:"username"`
}

// 获取data/config.ini里面的配置信息和版本信息
func AA_get_app_info(c *gin.Context) {
	var app AppInfo
	app.Version = cli.Version + "-" + cli.VersionDate
	app.Port = viper.GetString("servers.port")
	app.Runmode = viper.GetString("servers.runmode")
	app.Public_domain = config.Public_domain(c)
	app.Public_path = viper.GetString("storages.public_path")
	app.Username = viper.GetString("users.username")

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": app,
	})
	return
}

// 声明一个结构体，用于格式化服务器状态信息
type Server struct {
	OS        string                 `json:"os"`
	Arch      string                 `json:"arch"`
	CpuCount  int                    `json:"cpu_count"`
	DiskUsage *disk.UsageStat        `json:"disk_usage"`
	MemInfo   *mem.VirtualMemoryStat `json:"mem_info"`
	Load      *load.AvgStat          `json:"load"`
}

// 获取服务器状态信息
func AA_get_server_infos(c *gin.Context) {
	var server_info Server
	//获取操作系统
	server_info.OS = runtime.GOOS
	//获取架构
	server_info.Arch = runtime.GOARCH

	//获取CPU个数
	server_info.CpuCount, _ = cpu.Counts(true)
	//获取磁盘使用情况
	disk_usage, _ := disk.Usage("/")
	server_info.DiskUsage = disk_usage

	//获取内存信息
	server_info.MemInfo, _ = mem.VirtualMemory()
	//获取负载信息
	server_info.Load, _ = load.Avg()

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": server_info,
	})
	return
}
