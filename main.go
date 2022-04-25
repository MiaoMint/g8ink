package main

import (
	"github.com/beego/beego/v2/server/web/context"

	_ "g8ink/models"
	_ "g8ink/routers"
	"g8ink/tools"

	beego "github.com/beego/beego/v2/server/web"
)

var ADMIN_LOGIN_PASS, _ = beego.AppConfig.String("ADMIN_LOGIN_PASS")

func main() {

	//过滤未登录的
	var FilterUser = func(ctx *context.Context) {
		if ctx.Input.Session("Password") != ADMIN_LOGIN_PASS {
			ctx.Abort(404, "404")
		}
	}
	beego.InsertFilter("/admin/"+tools.GetAdminUrl()+"/api/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/admin/"+tools.GetAdminUrl()+"/home", beego.BeforeRouter, FilterUser)

	// orm.Debug = true

	beego.Run()
}
