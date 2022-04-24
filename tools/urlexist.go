package tools

import (
	"g8ink/models"

	"github.com/beego/beego/v2/client/orm"
)

func Urlexist(url string) string {
	o := orm.NewOrm()
	ob_url := models.Url{}
	err := o.QueryTable("url").Filter("OriginalUrl", url).One(&ob_url)
	if err != nil {
		return ""
	}
	return ob_url.ShortCode
}
