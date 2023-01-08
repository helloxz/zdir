package model

import (
	"fmt"
	"os"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// 声明全局变量
var (
	DB  *gorm.DB
	ERR error
)

// 初始化数据库连接
func init() {
	//如果数据库文件不存在，会自动创建
	DB, ERR = gorm.Open(sqlite.Open("data/db/zdir.db3"), &gorm.Config{})
	//如果出现错误，抛出错误并终止执行
	if ERR != nil {
		fmt.Println(ERR)
		os.Exit(1)
	} else {
		fmt.Print("Database connection succeeded!\n")
	}
}

// 导入默认数据库文件
func ImportDefaultSQL() {
	sql_file := "sql/init.sql"
	// 读取整个文件
	content, err := os.ReadFile(sql_file)

	//读取出现错误
	if err != nil {
		fmt.Println("read init.sql file failed, err:", err)
		return
	}

	//读取的数据转为字符串
	content_str := string(content)
	//sql文件按;拆分为sql语句
	rows := strings.Split(content_str, ";\n")

	//遍历执行SQL语句
	for _, sql := range rows {
		//如果SQL行是空的，则直接跳过
		if sql == "" {
			continue
		} else {
			//执行SQL语句
			DB.Exec(sql)
		}
	}
	fmt.Print("Initial data imported successfully\n")
}
