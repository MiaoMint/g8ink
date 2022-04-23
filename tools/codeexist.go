package tools

import "github.com/beego/beego/v2/client/orm"

func Codeexist(code string) bool {
	o := orm.NewOrm()
	return o.QueryTable("url").Filter("ShortCode", code).Exist()
}
