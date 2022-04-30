package tools

import (
	"context"
	"runtime"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/task"
)

type ipCache struct {
	Count    int
	Time     int64
	WaitTime int64
}

var li = make(map[string]ipCache)
var LIMIT_TIMES, _ = beego.AppConfig.Int("LIMIT_TIMES")
var Limit_Time, _ = beego.AppConfig.Int64("LIMIT_TIME")
var Limit_Wait_Time, _ = beego.AppConfig.Int64("LIMIT_WAIT_TIME")

func init() {
	logs.Info("初始化定时清理")
	var ipCacheClear = func(ctx context.Context) error {
		for k, ic := range li {
			// 判断当前时间距离创建的时间大于设定的时间并且被不是被惩罚用户时清理
			if time.Now().Unix()-ic.Time >= Limit_Time && time.Now().Unix() > ic.WaitTime {
				// logs.Info("删除缓存", ic.Time)
				delete(li, k)
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

	// 判断是否超过被罚时间超过清0
	if li[Ip].WaitTime != 0 && time.Now().Unix() > li[Ip].WaitTime {
		li[Ip] = ipCache{Count: 1}
		return false
	}

	// 判断是否在间隔时间内不在则清0
	if li[Ip].Time != 0 && time.Now().Unix()-li[Ip].Time >= Limit_Time {
		li[Ip] = ipCache{Count: 1, Time: 0, WaitTime: 0}
		return false
	}

	// 计数
	li[Ip] = ipCache{Count: li[Ip].Count + 1, Time: time.Now().Unix(), WaitTime: li[Ip].WaitTime}

	// 判断次数和判断被封禁
	if li[Ip].Count > LIMIT_TIMES || li[Ip].WaitTime != 0 {
		li[Ip] = ipCache{WaitTime: time.Now().Unix() + Limit_Wait_Time}
		return true
	}
	return false
}
