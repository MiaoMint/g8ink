package models

import (
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	// 注册驱动器
	orm.RegisterDriver("postgres", orm.DRPostgres)

	//获取环境变量 DSN
	DATABASE_URL, err := beego.AppConfig.String("DATABASE_URL")
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
	logs.Info(DATABASE_URL)
	// 注册数据库
	err = orm.RegisterDataBase("default", "postgres", DATABASE_URL)
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
	// 注册表
	orm.RegisterModel(new(Url), new(Ban))

	// 同步表
	orm.RunSyncdb("default", false, true)
}
