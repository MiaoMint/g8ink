/*
 * @Author: ERHECY 1981324730@qq.com
 * @Date: 2022-04-30 19:34:57
 * @LastEditors: ERHECY 1981324730@qq.com
 * @LastEditTime: 2022-06-03 01:40:46
 * @FilePath: \g8ink\tools\limitaccess.go
 * @Description:
 *
 * Copyright (c) 2022 by ERHECY 1981324730@qq.com, All Rights Reserved.
 */
package tools

import (
	"context"
	"runtime"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

type ipCache struct {
	Ip       string
	Count    int
	Time     int64
	WaitTime int64
}

var ip_cache = make(map[string]ipCache)

func init() {
	logs.Info("初始化定时清理")
	var ipCacheClear = func(ctx context.Context) error {
		for k, ic := range ip_cache {
			// 判断当前时间距离创建的时间大于设定的时间并且被不是被惩罚用户时清理
			if time.Now().Unix()-ic.Time >= LIMIT_TIME && time.Now().Unix() > ic.WaitTime {
				// logs.Info("删除缓存", ic.Time)
				delete(ip_cache, k)
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
	if ip_cache[Ip].WaitTime != 0 && time.Now().Unix() > ip_cache[Ip].WaitTime {
		ip_cache[Ip] = ipCache{Ip: Ip, Count: 1, Time: 0, WaitTime: 0}
		return false
	}

	// 判断是否在间隔时间内不在则清0
	if ip_cache[Ip].Time != 0 && time.Now().Unix()-ip_cache[Ip].Time >= LIMIT_TIME {
		ip_cache[Ip] = ipCache{Ip: Ip, Count: 1, Time: 0, WaitTime: 0}
		return false
	}

	// 计数
	ip_cache[Ip] = ipCache{Ip: Ip, Count: ip_cache[Ip].Count + 1, Time: time.Now().Unix(), WaitTime: ip_cache[Ip].WaitTime}

	// 判断次数和判断被封禁
	if ip_cache[Ip].Count > LIMIT_TIMES || ip_cache[Ip].WaitTime != 0 {
		ip_cache[Ip] = ipCache{Ip: Ip, WaitTime: time.Now().Unix() + LIMIT_WAIT_TIME}
		return true
	}
	return false
}

// 获取已经被限制的ip
func GetLimitIps() map[string]ipCache {
	limitip := make(map[string]ipCache)
	for _, ic := range ip_cache {
		if ic.Count > LIMIT_TIMES || ic.WaitTime != 0 {
			limitip[ic.Ip] = ic
		}
	}
	return limitip
}

// 解除被限制的ip
func DeleteLimitIp(Ip string) {
	ip_cache[Ip] = ipCache{Ip: Ip, Count: 1, Time: 0, WaitTime: 0}
}
