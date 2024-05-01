package tinyurl

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 函数接受一个字符串参数 url，并返回其MD5哈希值的十六进制字符串表示。
func md5Transform(url string) string {
	// 创建一个md5哈希对象
	hash := md5.New()
	// 将url写入哈希对象
	hash.Write([]byte(url))
	// 计算哈希值
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hashBytes)
}

//// CalculateTinyUrlPrefix 基于传入的md5使用随机hash计算string
//func CalculateTinyUrlPrefix(md5str string) string {
//
//}
