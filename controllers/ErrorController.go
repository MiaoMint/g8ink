package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "error404.html"
}

func (c *ErrorController) Error501() {
	c.Ctx.Output.Context.WriteString("501 error")
}
