package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func Load_ini() *ini.File {
	//载入配置文件
	cfg, err := ini.Load("data/config.ini")
	//如果载入配置出错，则终止执行
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}

// 返回公共存储的路径
func Public_path() string {
	//载入配置文件，通过cfg调用
	cfg := Load_ini()
	dir := cfg.Section("storages").Key("public_path").String()
	return dir
}

// 返回公共存储的域名
func Public_domain() string {
	//载入配置文件，通过cfg调用
	cfg := Load_ini()
	domain := cfg.Section("storages").Key("public_domain").String()
	return domain
}

// 返回端口
func Listen() string {
	//载入配置文件，通过cfg调用
	cfg := Load_ini()
	info := cfg.Section("servers").Key("port").String()
	return info
}

// 返回gin运行模式
func RunMode() string {
	//载入配置文件，通过cfg调用
	cfg := Load_ini()
	info := cfg.Section("servers").Key("RunMode").String()
	return info
}

// 返回站点信息
func Site_info() (string, string) {
	//载入配置文件，通过cfg调用
	cfg := Load_ini()
	title := cfg.Section("sites").Key("title").String()
	name := cfg.Section("sites").Key("name").String()

	return title, name
}
