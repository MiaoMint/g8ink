package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// 短链接对应表
type Url struct {
	Id int
	// 生成短链接后缀
	ShortCode string `orm:"index"`

	// 原链接
	OriginalUrl string `orm:"index"`

	// 生成时的ip
	Ip string `orm:"index"`

	// 生成时的时间
	Time time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Url) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

func Delete(Id string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("url").Filter("Id", Id).Delete()
	return err
}
