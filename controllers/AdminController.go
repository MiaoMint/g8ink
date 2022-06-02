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

// 短链接列表当前页面
var nowpage int

// 提示信息
var remessage string

// @router / [post,get]
func (c *AdminController) Login() {
	c.Data["title"] = "登录"

	// 判断是否登录 登录了302到后台首页
	if s, _ := c.GetSecureCookie(tools.GetCookiePass(), "Password"); s != tools.ADMIN_LOGIN_PASS {
		c.TplName = "admin/login.html"
	} else {
		c.Redirect("/admin/"+tools.GetAdminUrl()+"/home", 302)
	}

	// 登录请求
	if c.Ctx.Input.Method() == "POST" {
		if c.GetString("Password") == tools.ADMIN_LOGIN_PASS {
			c.SetSecureCookie(tools.GetCookiePass(), "Password", tools.ADMIN_LOGIN_PASS)
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
	page, _ := c.GetInt("page")
	nowpage = page

	o := orm.NewOrm()

	o.QueryTable("url").OrderBy("-Id").Limit(20, (page-1)*20).All(&url)
	c.Data["Linklist"] = &url

	c.Data["gnum"], _ = o.QueryTable("url").Count()

	//link列表页数
	c.Data["Linkpagenum"] = math.Ceil(float64(c.Data["gnum"].(int64)) / (float64)(20))

	// lnik列表当前页码
	c.Data["Linkpage"] = page

	// link列表下一页页码
	c.Data["Linknextpage"] = page + 1

	// link列表上一页页码
	c.Data["Linkpreviouspage"] = page - 1

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
	ban := []*models.Ban{}
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

// 白名单管理页面
// @router /whitelist [get]
func (c *AdminController) WhiteList() {
	c.Data["title"] = "临时封禁管理"
	c.Data["iswhitelist"] = 1

	c.Data["remessage"] = remessage

	o := orm.NewOrm()
	whitelist := []*models.WhiteList{}
	o.QueryTable("WhiteList").All(&whitelist)
	c.Data["WhiteList"] = &whitelist

	c.Layout = "admin/layout.html"
	c.TplName = "admin/whitelist.html"
	c.Data["Adminurl"] = tools.GetAdminUrl()

	remessage = ""
}

//删除link
// @router /api/DeleteLink [post,get]
func (c *AdminController) DeleteLink() {
	Id := c.GetString("id")
	err := models.Delete(Id)
	remessage = "删除成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/links?page="+strconv.Itoa(nowpage), 302)
}

//添加ban
// @router /api/AddBan [post,get]
func (c *AdminController) AddBan() {
	ban := models.Ban{
		Target: c.GetString("Target"),
		Type:   c.GetString("Type"),
	}
	_, err := ban.Insert()
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
	Id, _ := c.GetInt("id")
	ban := models.Ban{
		Id: Id,
	}
	_, err := ban.Delete()
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

// 添加白名单ip
// @router /api/AddWhiteList [post,get]
func (c *AdminController) AddWhiteList() {
	whitelist := models.WhiteList{
		Ip: c.GetString("ip"),
	}
	_, err := whitelist.Insert()
	remessage = "添加成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/whitelist", 302)
}

// 删除白名单ip
// @router /api/DeleteWhiteList [post,get]
func (c *AdminController) DeleteWhiteList() {
	Id, _ := c.GetInt("id")
	whitelist := models.WhiteList{
		Id: Id,
	}
	_, err := whitelist.Delete()
	remessage = "删除成功"
	if err != nil {
		remessage = err.Error()
	}
	c.Redirect("/admin/"+tools.GetAdminUrl()+"/whitelist", 302)
}
