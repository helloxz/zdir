package model

import "fmt"

// 声明一个结构体
type Z_login_log struct {
	ID         int    `gorm:"column:id"`
	Cid        string `gorm:"column:cid"`
	Token      string `gorm:"column:token"`
	Behavior   string `gorm:"column:behavior"`
	Expired_at int64  `gorm:"column:expired_at"`
	CreatedAt  int    `gorm:"column:created_at"`
	UpdatedAt  int    `gorm:"column:updated_at"`
	Ip         string `gorm:"column:ip"`
	Ua         string `gorm:"column:ua"`
	State      int    `gorm:"column:state"`
	Note       string `gorm:"column:note"`
}

// 插入登录记录,接收一个结构体，返回一个bool值
func LoginLogInsert(data Z_login_log) bool {
	//写入数据库日志表
	result := DB.Create(&data) // 通过数据的指针来创建

	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	} else {
		return true
	}
}
