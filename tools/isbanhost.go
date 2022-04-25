package tools

import "github.com/beego/beego/v2/client/orm"

//判断是否是被ban的域名
func Isbanhost(host string) bool {
	o := orm.NewOrm()
	return o.QueryTable("ban").Filter("Type", "host").Filter("Target", host).Exist()
}
