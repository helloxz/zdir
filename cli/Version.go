package cli

import "fmt"

var Version string
var VersionDate string

// 赋值全局变量
func init() {
	Version = "3.3.0"
	VersionDate = "20230326"
}

// 命令行打印版本
func GetVersion() {
	fmt.Printf(Version + "-" + VersionDate + "\n")
}
