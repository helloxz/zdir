package controller

import (
	"fmt"

	"github.com/coocood/freecache"
)

// 声明全局缓存参数
var Cache *freecache.Cache

// 缓存初始化
func init() {
	//设置一个最大为10M的缓存
	cacheSize := 100 * 1024 * 1024
	//初始化缓存
	Cache = freecache.NewCache(cacheSize)
}

// 缓存设置
func SetCache(key []byte, value []byte, ttl int) {
	//设置缓存
	Cache.Set(key, value, ttl)
}

// 获取缓存
func GetCache(key []byte) []byte {
	got, err := Cache.Get(key)
	if err != nil {
		fmt.Println(err)
		empty := []byte("")
		return empty
	} else {
		return got
	}
}

// 删除缓存
func DelCache(key string) bool {
	new_key := []byte(key)
	_ = Cache.Del(new_key)
	return true
}
