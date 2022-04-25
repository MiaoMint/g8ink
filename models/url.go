package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// 短链接对应表
type Url struct {
	Id int
	// 生成短链接后缀
	ShortCode string

	// 原链接
	OriginalUrl string

	// 生成时的ip
	Ip string

	// 生成时的时间s
	Time time.Time `orm:"auto_now_add;type(date)"`
}

func UrlInsert(shortcode string, originalurl string, ip string) error {
	o := orm.NewOrm()
	_, err := o.Insert(&Url{ShortCode: shortcode, OriginalUrl: originalurl, Ip: ip})
	return err
}

func UrlDelete(Id string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("url").Filter("Id", Id).Delete()
	return err
}
