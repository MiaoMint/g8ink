package controllers

import (
	"g8ink/models"

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
	if c.Ctx.GetCookie("Password") != ADMIN_LOGIN_PASS {
		c.Redirect("/", 302)
	}
	c.Layout = "admin/index.html"
	if c.Ctx.Input.Param(":path") == "home" {
		o := orm.NewOrm()
		c.Data["gnum"], _ = o.QueryTable("url").Count() //全站生成数量

		ban := []models.Ban{}
		o.QueryTable("ban").All(&ban)
		c.Data["Banlist"] = &ban
		c.TplName = "admin/home.html"
		return
	}
	c.Redirect("/", 302)
}

func (c *AdminController) AddBan() {
	Target := c.GetString("Target")
	Type := c.GetString("Type")
	// 判断是否登录
	if c.Ctx.GetCookie("Password") != ADMIN_LOGIN_PASS {
		c.Redirect("/", 302)
		return
	}
	if models.BanInsert(Type, Target) {
		c.Redirect("/admin/home?msg=添加成功#ban", 302)
	}
	c.Redirect("/admin/home?msg=添加错误#ban", 302)
}
