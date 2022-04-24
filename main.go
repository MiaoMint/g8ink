package main

import (
	_ "g8ink/models"
	_ "g8ink/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
