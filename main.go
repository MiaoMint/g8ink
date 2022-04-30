package main

import (
	"github.com/beego/beego/v2/server/web/context"

	_ "g8ink/models"
	_ "g8ink/routers"
	"g8ink/tools"

	beego "github.com/beego/beego/v2/server/web"
)

var ADMIN_LOGIN_PASS, _ = beego.AppConfig.String("ADMIN_LOGIN_PASS")
var Limit_Times, _ = beego.AppConfig.Int("Limit_Times")
var Limit_Time, _ = beego.AppConfig.Int64("Limit_Time")
var Limit_Wait_Time, _ = beego.AppConfig.Int64("Limit_Wait_Time")

func main() {

	//过滤未登录的
	var FilterUser = func(ctx *context.Context) {
		if ctx.Input.Session("Password") != ADMIN_LOGIN_PASS {
			ctx.Abort(404, "404")
		}
	}

	//过滤刷api的
	var FilterTimes = func(ctx *context.Context) {
		if ctx.Input.Method() == "POST" && tools.LimitAccess(ctx.Input.IP(), Limit_Times, Limit_Time, Limit_Wait_Time) {
			re := make(map[string]interface{})
			re["Code"] = -2
			re["Message"] = "太快啦~~要玩坏啦~＞︿＜"
			ctx.Output.JSON(&re, false, false)
		}
	}

	beego.InsertFilter("/admin/"+tools.GetAdminUrl()+"/api/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/admin/"+tools.GetAdminUrl()+"/home", beego.BeforeRouter, FilterUser)

	beego.InsertFilter("/", beego.BeforeRouter, FilterTimes)

	// orm.Debug = true

	beego.Run()
}
