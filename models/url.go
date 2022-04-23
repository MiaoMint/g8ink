package models

import "time"

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
