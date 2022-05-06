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

// @router / [post,get]
func (c *AdminController) Login() {
	c.Data["title"] = "登录"

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

// @router /home [get]
func (c *AdminController) Home() {
	c.Data["title"] = "后台管理"
	c.Data["remessage"] = remessage
	c.Data["ishome"] = 1
	c.Layout = "admin/layout.html"
	o := orm.NewOrm()

	//全站生成数量
	c.Data["gnum"], _ = o.QueryTable("url").Count()

	//获取被ban列表
	ban := []models.Ban{}
	o.QueryTable("ban").All(&ban)
	c.Data["Banlist"] = &ban

	//获取临时被ban列表
	c.Data["LimitIpList"] = tools.GetLimitIps()

	c.Data["Adminurl"] = tools.GetAdminUrl()

	c.TplName = "admin/home.html"
	remessage = ""
}

// 短链接管理页面
// @router /links [get]
func (c *AdminController) Links() {
	c.Data["title"] = "短链接管理"
	c.Data["islinks"] = 1

	url := []models.Url{}
	linkpage, _ := c.GetInt("page")

	if linkpage == 1 || linkpage == 0 {
		linkpage = 0
		nowpage = 1
	} else {
		linkpage = linkpage - 1
		nowpage = linkpage + 1
	}

	o := orm.NewOrm()

	o.QueryTable("url").OrderBy("-Id").Limit(20, linkpage*20).All(&url)
	c.Data["Linklist"] = &url

	c.Data["gnum"], _ = o.QueryTable("url").Count()

	//link列表页数
	c.Data["Linkpagenum"] = math.Ceil(float64(c.Data["gnum"].(int64)) / (float64)(20))

	// lnik列表当前页码
	c.Data["Linkpage"] = nowpage

	// link列表下一页页码
	c.Data["Linknextpage"] = nowpage + 1

	// link列表上一页页码
	c.Data["Linkpreviouspage"] = nowpage - 1

	c.Data["Adminurl"] = tools.GetAdminUrl()
	c.Data["remessage"] = remessage
	c.Layout = "admin/layout.html"
	c.TplName = "admin/links.html"
	remessage = ""
}

// 封禁管理页面
// @router /ban [get]
func (c *AdminController) Ban() {
	c.Data["title"] = "封禁管理"
	c.Data["isban"] = 1

	o := orm.NewOrm()
	c.Data["remessage"] = remessage

	//获取被ban列表
	ban := []models.Ban{}
	o.QueryTable("ban").All(&ban)
	c.Data["Banlist"] = &ban
	c.Layout = "admin/layout.html"
	c.TplName = "admin/ban.html"
	c.Data["Adminurl"] = tools.GetAdminUrl()

	remessage = ""

}

// 临时封禁管理页面
// @router /limitips [get]
func (c *AdminController) Limitips() {
	c.Data["title"] = "临时封禁管理"
	c.Data["islimitips"] = 1

	c.Data["remessage"] = remessage

	//获取临时被ban列表
	c.Data["LimitIpList"] = tools.GetLimitIps()

	c.Layout = "admin/layout.html"
	c.TplName = "admin/limitips.html"
	c.Data["Adminurl"] = tools.GetAdminUrl()

	remessage = ""

}

//删除link
// @router /api/DeleteLink [post,get]
func (c *AdminController) DeleteLink() {
	Id := c.GetString("id")
	err := models.UrlDelete(Id)
	remessage = "删除成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/links?page="+strconv.Itoa(nowpage), 302)
}

//添加ban
// @router /api/AddBan [post,get]
func (c *AdminController) AddBan() {
	Target := c.GetString("Target")
	Type := c.GetString("Type")
	err := models.BanInsert(Type, Target)
	remessage = "添加成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/ban", 302)
	// 更新banhost正则
	tools.GenerateRegularStr()
}

//删除ban
// @router /api/DeleteBan [post,get]
func (c *AdminController) DeleteBan() {
	Id := c.GetString("id")
	err := models.BanDelete(Id)
	remessage = "删除成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/ban", 302)
	// 更新banhost正则
	tools.GenerateRegularStr()
}

// 解除临时限制ip
// @router /api/DeleteLimitIp [post,get]
func (c *AdminController) DeleteLimitIp() {
	Ip := c.GetString("ip")
	tools.DeleteLimitIp(Ip)
	remessage = "解除限制成功"
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/limitips", 302)
}
