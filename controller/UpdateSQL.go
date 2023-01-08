package controller

import (
	"fmt"
	"os"
	"zdir/model"

	"github.com/gin-gonic/gin"
)

func UpdateSQL(c *gin.Context) {
	sql_name := c.Query("sql_name")
	//如果名字是空的，直接终止
	if sql_name == "" {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "SQL名称不能为空！",
			"data": "",
		})
		return
	}
	sql_file := "sql/" + sql_name
	//检测SQL文件是否存在
	if !V_is_file(sql_file) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "不是有效的SQL文件！",
			"data": "",
		})
		return
	}

	//执行SQL文件
	_, err := model.DbLogInsert(sql_name)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "SQL文件执行失败，请检查日志！",
			"data": "",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "",
		})
	}
}

// 创建数据库文件,当start的时候必定先执行此函数
func Create_db_file() {
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

	//连接数据库
	model.InitDB()

	//如果文件存在，则不再创建，但需要连接数据库
	if err == nil {
		fmt.Println("The database file already exists, skip this step.")
		return
	}

	//文件不存在，则创建文件并导入初始数据
	//导入默认的数据库文件，初始数据库的时候不存在zdir.db3会自动创建
	model.ImportDefaultSQL()
	fmt.Println("Database file created successfully!")
}
