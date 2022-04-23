package main

import (
	_ "g8url/models"
	_ "g8url/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	orm.Debug = true
	beego.Run()
}
