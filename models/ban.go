package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// ban表
type Ban struct {
	Id int
	// 被ban的类型
	Type string
	// 被ban的目标
	Target string
	// 生成时间
	Time time.Time `orm:"auto_now_add;type(date)"`
}

func BanInsert(Type string, Target string) error {
	o := orm.NewOrm()
	_, err := o.Insert(&Ban{Type: Type, Target: Target})
	return err
}

func BanDelete(Id string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("ban").Filter("Id", Id).Delete()
	return err
}
