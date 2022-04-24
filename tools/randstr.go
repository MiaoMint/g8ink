package tools

import (
	"math/rand"
	"time"
)

//生成随机字符串
func GetRandStr(length int) string {
	baseStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := 0; i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}

//生成短代码，判断是否生成相同的段代码了
//如果是则再次生成
func Getshortcode(length int) string {
	for {
		shortcode := GetRandStr(length)
		if !Codeexist(shortcode) {
			return shortcode
		}
	}
}
