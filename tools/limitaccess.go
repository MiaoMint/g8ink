package tools

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
)

type ipCache struct {
	Count    int
	Time     int64
	WaitTime int64
}

var li = make(map[string]ipCache)

// 限流访问 返回false代表许可 返回true代表禁止
func LimitAccess(Ip string, Times int, Time int64, WaitTime int64) bool {
	logs.Info(li)

	// 判断是否超过被罚时间超过清0
	if li[Ip].WaitTime != 0 && time.Now().Unix() > li[Ip].WaitTime {
		li[Ip] = ipCache{Count: 1}
		return false
	}

	// 判断是否在间隔时间内不在则清0
	if li[Ip].Time != 0 && time.Now().Unix()-li[Ip].Time >= Time {
		li[Ip] = ipCache{Count: 1, Time: 0, WaitTime: 0}
		return false
	}

	li[Ip] = ipCache{Count: li[Ip].Count + 1, Time: time.Now().Unix(), WaitTime: li[Ip].WaitTime}

	// 判断次数和判断被封禁
	if li[Ip].Count >= Times || li[Ip].WaitTime != 0 {
		li[Ip] = ipCache{WaitTime: time.Now().Unix() + WaitTime}
		return true
	}
	return false
}
