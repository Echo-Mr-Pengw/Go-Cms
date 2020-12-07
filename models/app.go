// 应用

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)
type App struct {
	Id uint `orm:"column(id); pk; auto; description(主键)"`
	Name string `orm:"column(name); size(64); default(); description(应用名)"`
	Status uint8 `orm:"column(status); default(1); description(状态 1:正常 2:冻结)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *App) TableEngine() string {
	return "INNODB"
}

func GetAllAppList() (*App, error) {
	app := new(App)
	_, err := orm.NewOrm().QueryTable(app).Filter("status", 1).All(app)
	return app, err
}