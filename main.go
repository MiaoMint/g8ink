package main

import (
	"github.com/beego/beego/v2/server/web/context"

	_ "g8ink/models"
	_ "g8ink/routers"

	beego "github.com/beego/beego/v2/server/web"
)

var ADMIN_LOGIN_PASS, _ = beego.AppConfig.String("ADMIN_LOGIN_PASS")

func main() {

	//过滤未登录的
	var FilterUser = func(ctx *context.Context) {
		if ctx.GetCookie("Password") != ADMIN_LOGIN_PASS {
			ctx.Redirect(301, "/")
		}
	}
	beego.InsertFilter("/admin/api/*", beego.BeforeRouter, FilterUser)

	beego.Run()
}
