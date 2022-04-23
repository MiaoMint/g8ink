package routers

import (
	"g8url/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")        //首页
	beego.Router("/:code", &controllers.MainController{}, "get:Get")   // 跳转
	beego.Router("/g", &controllers.MainController{}, "post:Generate") //生成短链接
}
