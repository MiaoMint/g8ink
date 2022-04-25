package tools

import beego "github.com/beego/beego/v2/server/web"

var ADMIN_URL, _ = beego.AppConfig.String("ADMIN_URL")

func GetAdminUrl() string {
	if ADMIN_URL == "unset" {
		ADMIN_URL = GetRandStr(6)
	}
	return ADMIN_URL
}
