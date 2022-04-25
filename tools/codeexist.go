package tools

import "github.com/beego/beego/v2/client/orm"

//判断Shortcode是否已经存在
func Codeexist(code string) bool {
	o := orm.NewOrm()
	return o.QueryTable("url").Filter("ShortCode", code).Exist()
}
