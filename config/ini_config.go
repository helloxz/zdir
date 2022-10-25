package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load_ini() {
	viper.SetConfigFile("data/config.ini") // 指定配置文件路径
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
	Load_ini()
	dir := viper.GetString("storages.public_path")
	return dir
}

// 返回公共存储的域名
func Public_domain() string {
	//载入配置文件，通过cfg调用
	Load_ini()
	domain := viper.GetString("storages.public_domain")
	return domain
}

// 返回端口
func Listen() string {
	//载入配置文件，通过cfg调用
	Load_ini()
	info := viper.GetString("servers.port")
	return info
}

// 返回gin运行模式
func RunMode() string {
	//载入配置文件，通过cfg调用
	Load_ini()
	info := viper.GetString("servers.RunMode")
	return info
}

// 返回站点信息
func Site_info() (string, string) {
	//载入配置文件，通过cfg调用
	Load_ini()
	title := viper.GetString("sites.title")
	name := viper.GetString("sites.name")

	return title, name
}

// 返回用户信息
func User_info() (string, string) {
	//载入配置文件，通过cfg调用
	Load_ini()
	username := viper.GetString("users.username")
	password := viper.GetString("users.password")

	return username, password
}
