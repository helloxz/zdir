// 验证各种参数是否合法
package controller

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 验证路径是否合法,返回一个BOOL值
func V_fpath(fpath string) bool {
	//判断用户传递的路径是否合法,如果存在.或..则返回false
	var validPath = regexp.MustCompile(`^(\.|\..).+`)
	v_re := validPath.MatchString(fpath)
	if v_re {
		return false
	} else if strings.Contains(fpath, "..") {
		return false
	} else if strings.Contains(fpath, "../") {
		return false
	} else if strings.Contains(fpath, "./") {
		return false
	} else {
		return true
	}
}

// 判断CID是否合法
func V_cid(cid string) bool {
	var valid = regexp.MustCompile(`^[a-zA-Z0-9]{6}$`)
	v_re := valid.MatchString(cid)

	if v_re {
		return true
	} else {
		return false
	}
}

// 判断路径是否是一个文件夹
func V_dir(dir string) bool {
	dirinfo, err := os.Stat(dir)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if dirinfo.IsDir() {
		return true
	} else {
		return false
	}
}

// 验证文件名，不应该包含/\|
func V_fname(fname string) bool {
	var valid = regexp.MustCompile(`(\/|\\|\|)+`)
	v_re := valid.MatchString(fname)
	if v_re {
		return false
	} else {
		return true
	}
}

// 验证是否是一个文件
func V_is_file(fpath string) bool {
	//获取文件信息
	finfo, err := os.Stat(fpath)

	//如果读取文件出现错误，比如不存在的情况，返回false
	if err != nil {
		return false
	} else {
		//如果是文件夹，返回false
		if finfo.IsDir() {
			return false
		} else {
			return true
		}
	}
}

// 验证搜索名称
func V_search_name(name string) bool {
	//正则验证，不能包含.. | & * exec --
	var valid = regexp.MustCompile(`(\.\.|\||&|\*|exec|--)+`)
	v_re := valid.MatchString(name)

	//name不能为空
	if name == "" {
		return false
	} else if len(name) > 24 {
		//不能超过24字节，英文字母占用1字节，中文占用3字节
		return false
	} else if v_re {
		return false
	} else {
		return true
	}
}
