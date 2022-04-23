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

func Getshortcode(length int) string {
	for {
		code := GetRandStr(6)
		if !Codeexist(code) {
			return code
		}
	}
}
