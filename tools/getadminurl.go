package tools

var cookiePass string

//获取后台链接
func GetAdminUrl() string {
	if ADMIN_URL == "unset" {
		ADMIN_URL = GetRandStr(6)
	}
	return ADMIN_URL
}

func GetCookiePass() string {
	if cookiePass == "" {
		cookiePass = GetRandStr(10)
	}
	return cookiePass
}
