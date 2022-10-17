package controller

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"zdir/config"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	//获取请求参数
	name := c.Query("name")

	//参数不能以.开头
	var validPath = regexp.MustCompile(`^(\.\.).+`)
	v_re := validPath.MatchString(name)
	if v_re {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "参数不合法！",
			"data": "",
		})
		c.Abort()
		return
	}

	//参数不能为空
	if name == "" {
		Err_json(1000, "参数不能为空", c)
		c.Abort()
		return
	}

	//载入配置文件，通过cfg调用
	public_dir := config.Public_path()

	//判断操作系统，非linux则不支持
	sysType := runtime.GOOS
	if sysType != "linux" {
		//返回json数据
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "非Linux服务端不支持全局搜索！",
			"data": "",
		})
		c.Abort()
		return
	}

	//参数中不能包含|或&或exec或*
	if strings.Contains(name, "|") || strings.Contains(name, "&") || strings.Contains(name, "exec") || strings.Contains(name, "*") {
		c.JSON(200, gin.H{
			"code": -1000,
			"msg":  "参数不合法！",
			"data": "",
		})
		c.Abort()
		return
	}

	//使用linux的find命令进行搜索,使用iname不区分大小写
	name = "*" + name + "*"
	out, _ := exec.Command("find", public_dir, "-type", "f", "-iname", name).Output()
	outstr := string(out)
	result := strings.Split(outstr, "\n")

	//声明文件类型
	var ftype string
	//声明一个结构体类型
	var new_info info
	//声明一个空的切片
	data := []info{}

	for _, value := range result {
		if value == "" {
			continue
		}
		//获取文件信息
		finfo, err := os.Stat(value)
		if err != nil {
			log.Fatal(err)
			return
		} else {
			//如果是目录
			if finfo.IsDir() {
				//文件类型赋值为folder
				ftype = "folder"
			} else {
				ftype = "file"
			}
			//替换字符串，去掉data/public
			value = strings.Replace(value, public_dir, "", -1)
			//继续获取其它信息
			new_info.Ftype = ftype
			new_info.Mtime = finfo.ModTime().Format("2006-01-02 15:04:05")
			new_info.Size = finfo.Size()
			name_temp := strings.Split(value, "/")
			new_info.Name = name_temp[len(name_temp)-1]
			new_info.Fpath = value
			//获取扩展名
			ext_temp := strings.Split(value, ".")
			//取分隔的最后一个元素
			new_info.Ext = strings.ToLower(ext_temp[len(ext_temp)-1])

			//追加到数据信息
			data = append(data, new_info)
		}
	}

	//返回json数据
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})

}

func Err_json(code int, msg string, c *gin.Context) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
}
