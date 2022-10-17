package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

// 命令行初始化
func Init() {
	//检查配置文件是否存在，如果存在了，则不进行初始化
	_, err := os.Stat("data/config.ini")
	//返回的error为空，说明文件存在，存在则不允许再次初始化
	if err == nil {
		fmt.Printf("Initialization failed, the configuration file already exists.\n")
		os.Exit(1)
	} else {
		//复制配置文件，参考：https://juejin.cn/post/6951352094003560484

		//打开源文件
		source, s_error := os.Open("config.simple.ini")
		if s_error != nil {
			fmt.Printf("s%\n", s_error)
			os.Exit(1)
		}

		//创建目标文件
		target, t_error := os.Create("data/config.ini")
		if t_error != nil {
			fmt.Printf("%s\n", t_error)
			os.Exit(1)
		}

		//关闭文件句柄
		defer source.Close()
		defer target.Close()

		//拷贝文件
		_, e := io.Copy(target, source)

		if e != nil {
			fmt.Printf("Failed to copy the configuration file, please check the permissions.\n")
			os.Exit(1)
		} else {
			//文件拷贝成功，继续执行初始化命令
			sysType := runtime.GOOS

			//注册windows服务
			if sysType == "windows" {
				windows_service()
				fmt.Printf("Init success.\n")
			} else if sysType == "linux" {
				linux_service()
				fmt.Printf("Init success.\n")
			} else {
				fmt.Printf("The current system does not support.\n")
				os.Exit(1)
			}
		}
	}
	//根据操作系统执行不同的命令
}

// linux添加服务
func linux_service() {
	_, err := exec.Command("bash", "sh/reg_service.sh").Output()

	if err != nil {
		fmt.Printf("Failed to register for the service\n")
		os.Exit(1)
	}
}

// windows添加服务
func windows_service() {
	//注册服务
	_, err1 := exec.Command("./run.exe", "install").Output()
	//运行服务
	//_, err2 := exec.Command("./run.exe", "start").Output()

	if err1 != nil {
		fmt.Printf("Failed to register for the service\n")
		os.Exit(1)
	}
}
