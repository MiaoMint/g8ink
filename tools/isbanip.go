package tools

import "github.com/beego/beego/v2/client/orm"

//判断是否是被ban的ip
func Isbanip(ip string) bool {
	o := orm.NewOrm()
	return o.QueryTable("ban").Filter("Type", "ip").Filter("Target", ip).Exist()
}
