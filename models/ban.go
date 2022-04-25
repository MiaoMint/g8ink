package models

import "github.com/beego/beego/v2/client/orm"

// ban表
type Ban struct {
	Id int
	// 被ban的类型
	Type string
	// 被ban的目标
	Target string
}

func BanInsert(Type string, Target string) bool {
	if (Type == "ip" || Type == "host") && Target != "" {
		o := orm.NewOrm()
		_, err := o.Insert(&Ban{Type: Type, Target: Target})
		return err == nil
	}
	return false
}
