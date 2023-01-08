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
	//配置文件路径
	config_file := "data/config/config.ini"
	//检查配置文件是否存在，如果存在了，则不进行初始化
	_, err := os.Stat(config_file)
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
		target, t_error := os.Create(config_file)
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

// 创建数据库文件，这个方法暂时没用了，改到UpdateSQL.go里面了
func create_db_file() {
	// 创建文件，在此之前检查目录是否存在
	// 检查文件夹是否存在
	db_dir := "data/db"

	if _, err := os.Stat(db_dir); os.IsNotExist(err) {
		// 创建文件夹
		err := os.Mkdir(db_dir, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//数据库文件路径
	db_file := "data/db/zdir.db3"
	//判断文件是否存在，存在就不再创建
	_, err := os.Stat(db_file)

	//如果文件存在，则不再创建
	if err == nil {
		fmt.Println("The database file already exists, skip this step.")
		return
	}

	//创建文件
	file, err := os.Create(db_file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("Database file created successfully!")
}

//导入默认数据库
