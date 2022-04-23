package tools

import "github.com/beego/beego/v2/client/orm"

func Isbanip(ip string) bool {
	o := orm.NewOrm()
	return o.QueryTable("ban").Filter("Type", "ip").Filter("Target", ip).Exist()
}
