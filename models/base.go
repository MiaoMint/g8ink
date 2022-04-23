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

	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "user=postgres password=root dbname=postgres host=127.0.0.1 port=5432 sslmode=disable"
	}
	// 注册数据库
	err := orm.RegisterDataBase("default", "postgres", dsn)
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
	// 注册表
	orm.RegisterModel(new(Url), new(Ban))

	// 生成表
	orm.RunSyncdb("default", false, true)
}
