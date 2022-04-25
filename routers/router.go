package routers

import (
	"g8ink/controllers"
	"g8ink/tools"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var ADMIN_URL, _ = beego.AppConfig.String("ADMIN_URL")
var HOST, _ = beego.AppConfig.String("HOST")

func init() {
	beego.ErrorController(&controllers.ErrorController{})             //错误处理
	beego.Router("/", &controllers.MainController{}, "get:Get")       //首页
	beego.Router("/", &controllers.MainController{}, "post:Generate") //生成短链接
	beego.Router("/:code", &controllers.MainController{}, "get:Get")  // 跳转

	if ADMIN_URL == "unset" {
		ADMIN_URL = tools.GetRandStr(6)
	}
	logs.Info("后台地址：", HOST+"/admin/"+ADMIN_URL)
	beego.Router("/admin/"+ADMIN_URL, &controllers.AdminController{}, "get:Login;post:Login") //后台登录
	beego.Router("/admin/home", &controllers.AdminController{}, "get:Home")                   //后台首页
	beego.Router("/admin/api/AddBan", &controllers.AdminController{}, "post:AddBan")          //添加ban
	beego.Router("/admin/api/DeleteBan", &controllers.AdminController{}, "get:DeleteBan")     //删除ban
	beego.Router("/admin/api/DeleteLink", &controllers.AdminController{}, "get:DeleteLink")   //删除lnik
}
