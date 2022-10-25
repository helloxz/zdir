package controller

import (
	"fmt"

	"github.com/coocood/freecache"
)

// 声明全局缓存参数
var Cache *freecache.Cache

// 缓存初始化
func InitCache() *freecache.Cache {
	//设置一个最大为10M的缓存
	cacheSize := 100 * 1024 * 1024
	//判断缓存是否存在
	if Cache != nil {
		return Cache
	} else {
		//初始化缓存
		Cache = freecache.NewCache(cacheSize)
		return Cache
	}
}

// 缓存设置
func SetCache(key []byte, value []byte, ttl int) {
	//设置缓存
	Cache.Set(key, value, ttl)
}

// 获取缓存
func GetCache(key []byte) []byte {
	Cache := InitCache()
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
	Cache := InitCache()
	new_key := []byte(key)
	_ = Cache.Del(new_key)
	return true
}
