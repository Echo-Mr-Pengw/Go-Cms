// 应用

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)
type App struct {
	Id uint `json:"id, "orm:"column(id); pk; auto; description(主键)"`
	Name string `json:"name, "orm:"column(name); size(64); default(); description(应用名), "form:"name"`
	Status uint8 `json:"status, "orm:"column(status); default(1); description(状态 1:正常 2:冻结), " form:"status"`
	CreateTime time.Time `json:"create_time, "orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *App) TableEngine() string {
	return "INNODB"
}

// 获取所有的应用
func GetAllAppList() ([]App, int64) {
	appList := []App{}
	num, _ := orm.NewOrm().QueryTable(new(App)).OrderBy("-id").All(&appList)
	return appList, num
}

// 添加应用
func AddAppInfo(app *App) (id int64) {
	id, _ = orm.NewOrm().Insert(app)
	return
}

// 通过应用ID获取各应用的信息
func GetDetailAppData(app *App) (err error) {
	err = orm.NewOrm().Read(app)
	return
}

// 更新应用信息
func UpdateAppData(app *App) (err error) {
	_, err = orm.NewOrm().Update(app);
	return
}