package main

import (
	_ "1-json/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Register driver
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// Register default database (creates file if not exists)
	orm.RegisterDataBase("default", "sqlite3", "data.db")

	web.BConfig.CopyRequestBody = true
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	web.Run()
}
