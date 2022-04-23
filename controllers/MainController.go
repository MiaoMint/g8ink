package controllers

import (
	"g8url/models"
	"g8url/tools"
	"os"
	"time"

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
		// 判断是否为链接是就跳转
		if len(url.OriginalUrl) > 7 && (url.OriginalUrl[0:7] == "http://" || url.OriginalUrl[0:8] == "https://") {
			// 跳转
			c.Redirect(url.OriginalUrl, 301)
			return
		}
		// 如果原url内容为空则跳转首页
		if url.OriginalUrl == "" {
			c.Redirect("", 301)
			return
		}
		c.Data["data"] = url.OriginalUrl
		c.TplName = "nourl.html"
		return
	}

	// 首页
	c.Data["gnum"], _ = o.QueryTable("url").Count()
	c.Data["unum"], _ = o.QueryTable("url").Filter("ip", c.Ctx.Input.IP()).Count()
	c.Data["year"] = time.Now().Year()
	c.TplName = "index.html"
}

func (c *MainController) Generate() {
	re := make(map[string]interface{})
	//获取表单数据
	shortcode := c.GetString("code")
	originalurl := c.GetString("url")

	if originalurl == "" || len(originalurl) > 500 || (len(shortcode) < 4 && len(shortcode) > 20) {
		re["Code"] = -1
		re["Message"] = "参数错误"
		c.Data["json"] = &re
		c.ServeJSON()
		return
	}

	//判断是否为封禁的ip
	if tools.Isbanip(c.Ctx.Input.IP()) {
		re["Code"] = -2
		re["Message"] = "你已被封禁"
		c.Data["json"] = &re
		c.ServeJSON()
		return
	}

	//判断是否为封禁的host
	if tools.Isbanhost(originalurl) {
		re["Code"] = -2
		re["Message"] = "封禁的域名"
		c.Data["json"] = &re
		c.ServeJSON()
		return
	}

	//生成code
	if shortcode == "" {
		shortcode = tools.Getshortcode(6)
	} else {
		//判断短代码是否存在
		if tools.Codeexist(shortcode) {
			re["Code"] = -1
			re["Message"] = "该短链接已存在"
			c.Data["json"] = &re
			c.ServeJSON()
			return
		}
	}

	//插入数据
	err := models.UrlInsert(shortcode, originalurl, c.Ctx.Input.IP())

	if err != nil {
		re["Code"] = -1
		re["Message"] = "生成错误"
	} else {
		re["Code"] = 200
		re["Shorturl"] = os.Getenv("HOST") + "/" + shortcode
		re["Message"] = "成功"
	}

	c.Data["json"] = &re
	c.ServeJSON()
}
