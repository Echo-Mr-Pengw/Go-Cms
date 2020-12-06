package main

import (
	_ "Permission-Platform/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"Permission-Platform/models"
)

func init() {

	mysqlDriver := beego.AppConfig.String("mysqldriver")
	mysqlHost   := beego.AppConfig.String("mysqlhost")
	mysqlUser   := beego.AppConfig.String("mysqluser")
	mysqlPass   := beego.AppConfig.String("mysqlpass")
	mysqlPort   := beego.AppConfig.String("mysqlport")
	dataBase    := beego.AppConfig.String("database")
	charset     := beego.AppConfig.String("charset")

	// root:root@tcp(127.0.0.1:3306)/perm_platform?charset=utf8
	dataSource := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort+")/" + dataBase + "?charset=" + charset

	// 注册model
	orm.RegisterModelWithPrefix("tbl_", new(models.Admin))
	// 注册驱动
	orm.RegisterDriver(mysqlDriver, orm.DRMySQL)
	// 注册db
	orm.RegisterDataBase("default", mysqlDriver, dataSource)

	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}

