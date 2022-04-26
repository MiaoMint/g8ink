package controllers

import (
	"g8ink/models"
	"g8ink/tools"
	"math"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminController struct {
	beego.Controller
}

var ADMIN_LOGIN_PASS, _ = beego.AppConfig.String("ADMIN_LOGIN_PASS")

// 短链接列表当前页面
var nowpage int

// 提示信息
var remessage string

func (c *AdminController) Login() {
	c.Data["title"] = "登录"
	c.Layout = "admin/layout.html"
	// 判断是否登录 登录了302到后台首页
	if c.GetSession("Password") != ADMIN_LOGIN_PASS {
		c.TplName = "admin/login.html"
	} else {
		c.Redirect("/admin/"+tools.GetAdminUrl()+"/home", 302)
	}

	// 登录请求
	if c.Ctx.Input.Method() == "POST" {
		if c.GetString("Password") == ADMIN_LOGIN_PASS {
			c.SetSession("Password", ADMIN_LOGIN_PASS)
			c.Redirect("/admin/"+tools.GetAdminUrl()+"/home", 302)
		} else {
			c.Data["remessage"] = "密码错误"
		}
	}
}

func (c *AdminController) Home() {
	c.Data["title"] = "后台管理"
	c.Data["remessage"] = remessage
	c.Layout = "admin/layout.html"
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

	if linkpage == 1 || linkpage == 0 {
		linkpage = 0
		nowpage = 1
	} else {
		linkpage = linkpage - 1
		nowpage = linkpage + 1
	}

	o.QueryTable("url").OrderBy("-Id").Limit(20, linkpage*20).All(&url)
	c.Data["Linklist"] = &url

	//link列表页数
	c.Data["Linkpagenum"] = math.Ceil(float64(c.Data["gnum"].(int64)) / (float64)(20))

	// lnik列表当前页码
	c.Data["Linkpage"] = nowpage

	// link列表下一页页码
	c.Data["Linknextpage"] = nowpage + 1

	// link列表上一页页码
	c.Data["Linkpreviouspage"] = nowpage - 1

	c.Data["Adminurl"] = tools.GetAdminUrl()

	c.TplName = "admin/home.html"
	remessage = ""
}

//删除link
func (c *AdminController) DeleteLink() {
	Id := c.GetString("id")
	err := models.UrlDelete(Id)
	remessage = "删除成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/home?linkpage="+strconv.Itoa(nowpage)+"#link", 302)
}

//添加ban
func (c *AdminController) AddBan() {
	Target := c.GetString("Target")
	Type := c.GetString("Type")
	err := models.BanInsert(Type, Target)
	remessage = "添加成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/home#ban", 302)
}

//删除ban
func (c *AdminController) DeleteBan() {
	Id := c.GetString("id")
	err := models.BanDelete(Id)
	remessage = "删除成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/home#ban", 302)
}
