package controller

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"zdir/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// 声明一个全局变量用来获取.ini配置
var cfg *ini.File

// 定义一个结构体，用来存放文件或文件夹信息
type info struct {
	Name  string
	Size  int64
	Mtime string
	Ftype string
	Fpath string
	Ext   string
	Link  string
}

func FileList(c *gin.Context) {
	//获取公共存储的路径
	public_dir := config.Public_path()
	storage_domain := config.Public_domain()
	//获取请求参数
	path := c.Query("path")
	//判断用户传递的路径是否合法
	var validPath = regexp.MustCompile(`^(\.|\..).+`)
	v_re := validPath.MatchString(path)
	if v_re {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件夹不合法！",
			"data": "",
		})
		c.Abort()
		return
	}
	//组合完整路径
	var full_path string
	if path == "" {
		full_path = public_dir
	} else {
		full_path = public_dir + "/" + path
	}
	//判断文件是否存在，如果不存在，则终止执行
	_, err := os.Stat(full_path)
	if os.IsNotExist(err) {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "文件夹不存在！",
			"data": "",
		})
		c.Abort()
		return
	}

	//声明文件类型
	var ftype string
	//声明一个结构体类型
	var new_info info
	//声明一个空的切片
	result := []info{}
	files, err := ioutil.ReadDir(full_path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//获取文件或文件夹路径
		fpath := full_path + "/" + file.Name()
		fname := file.Name()
		finfo, err := os.Stat(fpath)
		if err != nil {
			log.Fatal(err)
			return
		} else {
			//如果是隐藏文件，直接跳过(.开头的视为隐藏文件)
			var validHide = regexp.MustCompile(`^\..*`)
			v_re := validHide.MatchString(fname)
			if v_re {
				continue
			}

			//如果是目录
			if finfo.IsDir() {
				//文件类型赋值为folder
				ftype = "folder"
				new_info.Ext = ""
				new_info.Link = ""
			} else {
				ftype = "file"
				//获取扩展名
				ext_temp := strings.Split(fname, ".")
				//取分隔的最后一个元素
				new_info.Ext = strings.ToLower(ext_temp[len(ext_temp)-1])
			}
			//继续获取其它信息
			new_info.Ftype = ftype
			new_info.Mtime = finfo.ModTime().Format("2006-01-02 15:04:05")
			new_info.Size = finfo.Size()
			new_info.Name = fname
			new_info.Fpath = path + "/" + fname
			new_info.Link = storage_domain + "/public" + url.QueryEscape(new_info.Fpath)

			//追加到数据信息
			result = append(result, new_info)
		}

	}
	//返回json数据
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": result,
	})
}
