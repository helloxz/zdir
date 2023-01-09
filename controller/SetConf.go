// 设置配置文件
package controller

import (
	"fmt"

	"github.com/spf13/viper"
)

// 设置配置文件的键值
func SetKVS(Key string, Value string) bool {
	//判断key是否存在,不存在，则设置
	viper.Set(Key, Value)
	//viper.SetConfigType("toml") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	err := viper.WriteConfig()

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}
