package tinyurl

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// md5 函数接受一个字符串参数 url，并返回其MD5哈希值的十六进制字符串表示。
func Md5Transform(url string) string {
	// 创建一个md5哈希对象
	hash := md5.New()
	// 将url写入哈希对象
	hash.Write([]byte(url))
	// 计算哈希值
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hashBytes)
}

const (
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	charLen = uint64(len(chars))
)

// CalculateTinyUrlPrefix 生成短链接前缀, 通过randFix加盐，在出现
func CalculateTinyUrlPrefix(md5str string, randFix uint64) string {
	// 计算传入字符串的MD5哈希值
	hash := md5.Sum([]byte(md5str))

	// 创建一个短链接前缀
	shortURLPrefix := make([]byte, 6) // 假设我们想要一个6字符长度的短链接

	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	for i := range shortURLPrefix {
		// 将hash的字节作为索引，并转换为uint64类型
		byteIndex := uint64(hash[i%md5.Size])
		// 使用模运算找到chars中对应的字符索引, 对byteIndex进行加盐处理
		charIndex := int((byteIndex + randFix) % charLen)
		shortURLPrefix[i] = chars[charIndex]
	}

	return string(shortURLPrefix)
}

func main() {
	// 示例字符串
	md5Str := "https://kimi.moonshot.cn/chat/covpersubmsb8l3m9hlg?data_source=tracer&utm_campaign=TR_PbzLg3eV&utm_content=&utm_medium=%E5%BE%AE%E8%BD%AFbing&utm_source=bing&utm_term=&msclkid=0775a2f498a7128ca080c4cb0fecc169"
	start := time.Now().UnixMilli()
	md5Str = Md5Transform(md5Str)
	// 计算MD5哈希值
	md5Hash := fmt.Sprintf("%x", md5.Sum([]byte(md5Str)))

	// 生成短链接前缀
	shortURL := CalculateTinyUrlPrefix(md5Hash, 0)
	end := time.Now().UnixMilli()
	fmt.Println("MD5 Hash:", md5Hash)
	fmt.Println("Short URL:", shortURL)
	fmt.Println("time:", end-start)
}
