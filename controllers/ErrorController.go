/*
 * @Author: ERHECY 1981324730@qq.com
 * @Date: 2022-04-24 07:40:06
 * @LastEditors: ERHECY 1981324730@qq.com
 * @LastEditTime: 2022-06-12 17:32:13
 * @FilePath: \g8ink\controllers\ErrorController.go
 * @Description: 错误处理
 *
 * Copyright (c) 2022 by ERHECY 1981324730@qq.com, All Rights Reserved.
 */
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
	c.Ctx.WriteString("501 error")
}
