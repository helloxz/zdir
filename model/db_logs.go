package model

import (
	"fmt"
	"os"
	"strings"
)

// 声明一个结构体
type z_db_logs struct {
	ID        int    `gorm:"column:id"`
	Sql_name  string `gorm:"column:sql_name"`
	State     int    `gorm:"column:state"`
	Note      string `gorm:"column:note"`
	CreatedAt int    `gorm:"column:created_at"`
	UpdatedAt int    `gorm:"column:updated_at"`
}

func DbLogInsert(sql_name string) (int64, error) {
	log := z_db_logs{
		Sql_name: sql_name,
		State:    1,
	}

	sql_file := "sql/" + sql_name
	// 读取整个文件
	content, err := os.ReadFile(sql_file)

	//读取出现错误
	if err != nil {
		fmt.Println("read init.sql file failed, err:", err)
		return 0, err
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

	//写入数据库日志表
	result := DB.Create(&log) // 通过数据的指针来创建
	return result.RowsAffected, result.Error
}
