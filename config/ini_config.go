package config

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	//默认配置文件
	default_config_file := "data/config/confi.ini"
	//备用配置文件
	backup_config_file := "data/config.ini"
	var config_file string

	//首先读取默认配置文件，如果不存在，则读取备用配置文件
	if v_is_file(default_config_file) {
		config_file = default_config_file
	} else {
		config_file = backup_config_file
	}

	viper.SetConfigFile(config_file) // 指定配置文件路径
	//指定ini类型的文件
	viper.SetConfigType("ini")
	err := viper.ReadInConfig() // 读取配置信息

	if err != nil { // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// 返回公共存储的路径
func Public_path() string {
	//载入配置文件，通过cfg调用
	dir := viper.GetString("storages.public_path")
	return dir
}

// 返回公共存储的域名
func Public_domain(c *gin.Context) string {
	//获取请求host
	host := c.Request.Host
	//载入配置文件，通过cfg调用
	domain := viper.GetString("storages.public_domain")
	//如果公共存储域名是空的，则加上public
	if domain == "" {
		domain = "http://" + host + "/public"
	}
	return domain
}

// 返回端口
func Listen() string {
	//载入配置文件，通过cfg调用
	info := viper.GetString("servers.port")
	return info
}

// 返回gin运行模式
func RunMode() string {
	//载入配置文件，通过cfg调用
	info := viper.GetString("servers.RunMode")
	return info
}

// 返回站点信息
func Site_info() (string, string) {
	//载入配置文件，通过cfg调用
	title := viper.GetString("sites.title")
	name := viper.GetString("sites.name")

	return title, name
}

// 返回用户信息
func User_info() (string, string) {
	//载入配置文件，通过cfg调用
	username := viper.GetString("users.username")
	password := viper.GetString("users.password")

	return username, password
}

func Base64(str string) string {
	encodeStr := base64.StdEncoding.EncodeToString([]byte(str))
	return encodeStr
}

// 验证是否是一个文件
func v_is_file(fpath string) bool {
	//获取文件信息
	finfo, err := os.Stat(fpath)

	//如果读取文件出现错误，比如不存在的情况，返回false
	if err != nil {
		return false
	} else {
		//如果是文件夹，返回false
		if finfo.IsDir() {
			return false
		} else {
			return true
		}
	}
}
