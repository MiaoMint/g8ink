package tools

import (
	"g8ink/models"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

var RegularStr string //生成的正则表达式

func init() {
	GenerateRegularStr()
}

//判断是否是被ban的域名
func Isbanhost(host string) bool {
	valid := validation.Validation{}
	// 为空使用正则
	if RegularStr == "" {
		return false
	}
	return valid.Match(host, regexp.MustCompile(RegularStr), "url").Ok
}

// 生成正则表达式
func GenerateRegularStr() {

	// 清空正则
	RegularStr = ""

	banlist := []models.Ban{}
	o := orm.NewOrm()
	o.QueryTable("ban").Filter("Type", "host").All(&banlist)

	// 拼接表达式
	for i, b := range banlist {
		if i == len(banlist)-1 {
			RegularStr += "(" + strings.Replace(b.Target, ".", `\.`, -1) + ")"
			break
		}
		RegularStr += "(" + strings.Replace(b.Target, ".", `\.`, -1) + ")|"
	}

	logs.Info("更新BanHost正则", RegularStr)
}
