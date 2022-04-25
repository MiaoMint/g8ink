package controllers

import (
	"g8ink/models"
	"math"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminController struct {
	beego.Controller
}

var ADMIN_LOGIN_PASS, _ = beego.AppConfig.String("ADMIN_LOGIN_PASS")

func (c *AdminController) Login() {
	c.Data["title"] = "登录"
	c.Layout = "admin/index.html"
	// 判断是否登录 登录了302到后台首页
	if c.Ctx.GetCookie("Password") != ADMIN_LOGIN_PASS {
		c.TplName = "admin/login.html"
	} else {
		c.Redirect("/admin/home", 302)
	}

	// 登录请求
	if c.Ctx.Input.Method() == "POST" {
		if c.GetString("Password") == ADMIN_LOGIN_PASS {
			c.Ctx.SetCookie("Password", ADMIN_LOGIN_PASS)
			c.Redirect("/admin/home", 302)
		} else {
			c.Data["remessage"] = "密码错误"
		}
	}
}

func (c *AdminController) Home() {
	c.Data["title"] = "后台管理"
	c.Data["remessage"] = c.GetString("msg")
	c.Layout = "admin/index.html"
	o := orm.NewOrm()

	//全站生成数量
	c.Data["gnum"], _ = o.QueryTable("url").Count()

	//获取被ban列表
	ban := []models.Ban{}
	o.QueryTable("ban").All(&ban)
	c.Data["Banlist"] = &ban

	//获取link列表
	url := []models.Url{}
	linkpage, _ := c.GetInt("linkpage")
	if linkpage == 0 {
		linkpage = 1
	}
	o.QueryTable("url").Limit(30, linkpage*30).All(&url)
	c.Data["Linklist"] = &url

	//link列表页数
	c.Data["Linkpagenum"] = math.Ceil(float64(c.Data["gnum"].(int64)) / (float64)(30))

	// lnik列表当前页码
	c.Data["Linkpage"] = linkpage

	// link列表下一页页码
	c.Data["Linknextpage"] = linkpage + 1

	// link列表上一页页码
	c.Data["Linkpreviouspage"] = linkpage - 1

	c.TplName = "admin/home.html"
}

//删除link
func (c *AdminController) DeleteLink() {
	Id := c.GetString("id")
	err := models.UrlDelete(Id)
	if err != nil {
		c.Redirect("/admin/home?msg="+err.Error()+"#link", 302)
	}
	c.Redirect("/admin/home?msg=删除成功#link", 302)
}

//添加ban
func (c *AdminController) AddBan() {
	Target := c.GetString("Target")
	Type := c.GetString("Type")
	err := models.BanInsert(Type, Target)
	if err != nil {
		c.Redirect("/admin/home?msg="+err.Error()+"#ban", 302)
	}
	c.Redirect("/admin/home?msg=添加成功#ban", 302)
}

//删除ban
func (c *AdminController) DeleteBan() {
	Id := c.GetString("id")
	err := models.BanDelete(Id)
	if err != nil {
		c.Redirect("/admin/home?msg="+err.Error()+"#ban", 302)
	}
	c.Redirect("/admin/home?msg=删除成功#ban", 302)
}
