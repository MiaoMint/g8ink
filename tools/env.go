/*
 * @Author: ERHECY 1981324730@qq.com
 * @Date: 2022-06-03 01:40:19
 * @LastEditors: ERHECY 1981324730@qq.com
 * @LastEditTime: 2022-06-03 02:04:38
 * @FilePath: \g8ink\tools\env.go
 * @Description:
 *
 * Copyright (c) 2022 by ERHECY 1981324730@qq.com, All Rights Reserved.
 */
package tools

import beego "github.com/beego/beego/v2/server/web"

var LIMIT_TIMES, _ = beego.AppConfig.Int("LIMIT_TIMES")
var LIMIT_TIME, _ = beego.AppConfig.Int64("LIMIT_TIME")
var LIMIT_WAIT_TIME, _ = beego.AppConfig.Int64("LIMIT_WAIT_TIME")

var ADMIN_URL, _ = beego.AppConfig.String("ADMIN_URL")
var ADMIN_LOGIN_PASS, _ = beego.AppConfig.String("ADMIN_LOGIN_PASS")

var HOST, _ = beego.AppConfig.String("HOST")
var MAX_URL, _ = beego.AppConfig.Int("MAX_URL")
var MAX_SHORTCODE, _ = beego.AppConfig.Int("MAX_SHORTCODE")
var MIN_SHORTCODE, _ = beego.AppConfig.Int("MIN_SHORTCODE")
var RAND_SHORTCODE, _ = beego.AppConfig.Int("RAND_SHORTCODE")

var WEB_SCRIPT, _ = beego.AppConfig.String("WEB_SCRIPT")
var WEB_BACKGROUND, _ = beego.AppConfig.String("WEB_BACKGROUND")
