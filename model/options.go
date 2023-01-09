// z_options模型
package model

import (
	"encoding/json"
	"fmt"
)

// 声明一个结构体
type z_option struct {
	ID        int    `gorm:"column:id"`
	Key       string `gorm:"column:key"`
	Value     string `gorm:"column:value"`
	Note      string `gorm:"column:note"`
	CreatedAt int    `gorm:"column:created_at"`
	UpdatedAt int    `gorm:"column:updated_at"`
}

// 设置key->value
func OptionSet(key string, value string, note ...string) (int64, error) {
	var note_content string
	//判断可变参数note是否存在,如果存在，则赋值给note_content
	if len(note) > 0 {
		note_content = note[0]
	}

	option := z_option{Key: key, Value: value, Note: note_content}
	//在插入之前检测key是否存在，存在则update，不存在则insert
	q_result := DB.Where("key = ?", key).First(&option)

	//如果影响行数为1，说明个存在数据，则应该更新
	if q_result.RowsAffected == 1 {
		fmt.Println("更新操作！")
		//更新单个例
		result := DB.Model(&z_option{}).Where("key = ?", key).Update("value", value)
		return result.RowsAffected, result.Error
	} else {
		//插入数据
		result := DB.Create(&option) // 通过数据的指针来创建
		return result.RowsAffected, result.Error
	}

}

// 通过key获取数据
func OptionGet(key string) map[string]interface{} {
	option := z_option{}
	q_result := DB.Where("key = ?", key).First(&option)
	if q_result.Error != nil {
		fmt.Println(q_result.Error)
		var empty map[string]interface{}
		return empty
	}
	//声明一个接口
	var v interface{}
	//声明一个[]byte变量并赋值
	value := []byte(option.Value)
	//字符串转json，注意第一个变量必须是[]byte类型，第二个变量为接口指针
	json.Unmarshal(value, &v)
	//赋值给变量，变量类型是：map[string]interface{}
	data := v.(map[string]interface{})
	return data
}
