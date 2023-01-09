package cli

import "fmt"

var Version string
var VersionDate string

// 赋值全局变量
func init() {
	Version = "3.2.0"
	VersionDate = "20230109"
}

// 命令行打印版本
func GetVersion() {
	fmt.Printf(Version + "-" + VersionDate + "\n")
}
