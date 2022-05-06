package routers

import (
	"g8ink/controllers"
	"g8ink/tools"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var ADMIN_URL = tools.GetAdminUrl()
var HOST, _ = beego.AppConfig.String("HOST")

func init() {
	beego.ErrorController(&controllers.ErrorController{})                      //错误处理
	beego.Router("/", &controllers.MainController{}, "get:Home;post:Generate") //首页,生成短链接
	beego.Router("/:code", &controllers.MainController{}, "get:Home")          // 跳转
	beego.Router("/robots.txt", &controllers.MainController{}, "get:Robots")   // robots.txt

	// 提示信息
	logs.Info("首页地址：", HOST)
	logs.Info("后台地址：", HOST+"/admin/"+ADMIN_URL)

	ns := beego.NewNamespace("/admin/"+ADMIN_URL,
		beego.NSInclude(
			&controllers.AdminController{},
		),
	)
	beego.AddNamespace(ns)
}
