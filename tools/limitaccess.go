package tools

import (
	"context"
	"runtime"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/task"
)

type ipCache struct {
	Ip       string
	Count    int
	Time     int64
	WaitTime int64
}

var IP_CACHE = make(map[string]ipCache)
var LIMIT_TIMES, _ = beego.AppConfig.Int("LIMIT_TIMES")
var LIMIT_TIME, _ = beego.AppConfig.Int64("LIMIT_TIME")
var LIMIT_WAIT_TIME, _ = beego.AppConfig.Int64("LIMIT_WAIT_TIME")

func init() {
	logs.Info("初始化定时清理")
	var ipCacheClear = func(ctx context.Context) error {
		for k, ic := range IP_CACHE {
			// 判断当前时间距离创建的时间大于设定的时间并且被不是被惩罚用户时清理
			if time.Now().Unix()-ic.Time >= LIMIT_TIME && time.Now().Unix() > ic.WaitTime {
				// logs.Info("删除缓存", ic.Time)
				delete(IP_CACHE, k)
			}
		}
		runtime.GC()
		return nil
	}
	tk := task.NewTask("ipCacheClear", "0 15 03 * * *", ipCacheClear)
	task.AddTask("ipCacheClear", tk)
	task.StartTask()
}

// 限流访问 返回false代表许可 返回true代表禁止
func LimitAccess(Ip string) bool {
	if orm.NewOrm().QueryTable("WhiteList").Filter("Ip", Ip).Exist() {
		return false
	}
	// 判断是否超过被罚时间超过清0
	if IP_CACHE[Ip].WaitTime != 0 && time.Now().Unix() > IP_CACHE[Ip].WaitTime {
		IP_CACHE[Ip] = ipCache{Ip: Ip, Count: 1, Time: 0, WaitTime: 0}
		return false
	}

	// 判断是否在间隔时间内不在则清0
	if IP_CACHE[Ip].Time != 0 && time.Now().Unix()-IP_CACHE[Ip].Time >= LIMIT_TIME {
		IP_CACHE[Ip] = ipCache{Ip: Ip, Count: 1, Time: 0, WaitTime: 0}
		return false
	}

	// 计数
	IP_CACHE[Ip] = ipCache{Ip: Ip, Count: IP_CACHE[Ip].Count + 1, Time: time.Now().Unix(), WaitTime: IP_CACHE[Ip].WaitTime}

	// 判断次数和判断被封禁
	if IP_CACHE[Ip].Count > LIMIT_TIMES || IP_CACHE[Ip].WaitTime != 0 {
		IP_CACHE[Ip] = ipCache{Ip: Ip, WaitTime: time.Now().Unix() + LIMIT_WAIT_TIME}
		return true
	}
	return false
}

// 获取已经被限制的ip
func GetLimitIps() map[string]ipCache {
	limitip := make(map[string]ipCache)
	for _, ic := range IP_CACHE {
		if ic.Count > LIMIT_TIMES || ic.WaitTime != 0 {
			limitip[ic.Ip] = ic
		}
	}
	return limitip
}

// 解除被限制的ip
func DeleteLimitIp(Ip string) {
	IP_CACHE[Ip] = ipCache{Ip: Ip, Count: 1, Time: 0, WaitTime: 0}
}
