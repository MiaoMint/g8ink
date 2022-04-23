package tools

import "github.com/beego/beego/v2/client/orm"

func Isbanhost(host string) bool {
	o := orm.NewOrm()
	return o.QueryTable("ban").Filter("Type", "host").Filter("Target", host).Exist()
}
