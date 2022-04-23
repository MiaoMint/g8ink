package main

import (
	_ "g8url/models"
	_ "g8url/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
