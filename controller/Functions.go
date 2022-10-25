// 常用方法集合
package controller

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

// md5字符串加密
func md5s(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成一个随机字符串
func RandStr(n int) string {
	var bytes []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Int31()%62]
	}
	return string(result)
}
