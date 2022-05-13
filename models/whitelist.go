package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type WhiteList struct {
	Id   int
	Ip   string
	Time time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *WhiteList) Insert() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

func (m *WhiteList) Delete() (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}
