package models

import (
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/lib/pq"
)

func init() {
	// 注册驱动器
	orm.RegisterDriver("postgres", orm.DRPostgres)

	//获取环境变量 DSN
	DATABASE_URL := os.Getenv("DATABASE_URL")
	logs.Info(DATABASE_URL)
	if DATABASE_URL == "" {
		DATABASE_URL = "user=postgres password=root dbname=postgres host=127.0.0.1 port=5432 sslmode=disable"
	}
	// 注册数据库
	err := orm.RegisterDataBase("default", "postgres", DATABASE_URL)
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
	// 注册表
	orm.RegisterModel(new(Url), new(Ban))

	// 同步表
	orm.RunSyncdb("default", false, true)
}
