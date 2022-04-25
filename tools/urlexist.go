package tools

import (
	"g8ink/models"

	"github.com/beego/beego/v2/client/orm"
)

//判断url是否已存在
func Urlexist(url string) string {
	o := orm.NewOrm()
	ob_url := models.Url{}
	err := o.QueryTable("url").Filter("OriginalUrl", url).One(&ob_url)
	if err != nil {
		return ""
	}
	return ob_url.ShortCode
}
