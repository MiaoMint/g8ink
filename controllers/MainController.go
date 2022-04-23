package controllers

import (
	"g8url/models"
	"g8url/tools"
	"os"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	code := c.Ctx.Input.Param(":code")
	o := orm.NewOrm()

	//如果shortcode不等于空则在数据库里找是否存在
	if code != "" {
		url := models.Url{}
		o.QueryTable("url").Filter("ShortCode", code).One(&url)
		c.Redirect(url.OriginalUrl, 301)
	}

	c.Data["gnum"], _ = o.QueryTable("url").Count()
	c.TplName = "index.html"
}

func (c *MainController) Generate() {
	re := make(map[string]interface{})
	//获取表单数据
	shortcode := c.GetString("code")
	originalurl := c.GetString("url")

	if originalurl == "" {
		re["Code"] = -1
		re["Message"] = "参数错误"
		c.Data["json"] = &re
		c.ServeJSON()
		return
	}

	if shortcode == "" {
		shortcode = tools.GetRandStr(6)
	}

	//插入数据库
	o := orm.NewOrm()
	url := models.Url{ShortCode: shortcode, OriginalUrl: originalurl, Ip: c.Ctx.Input.IP()}
	_, err := o.Insert(&url)
	if err != nil {
		re["Code"] = -1
		re["Message"] = "生成错误"
	} else {
		re["Code"] = 200
		re["Shorturl"] = os.Getenv("HOST") + shortcode
		re["Message"] = "成功"
	}

	c.Data["json"] = &re
	c.ServeJSON()
}
