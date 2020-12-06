package models

import (
	"time"
)

type Admin struct {
	Id uint `orm:"column(id); pk; auto; description(主键)"`
	Emplid string `orm:"column(emplid); size(32); default(); description(管理员工号)"`
	Name string `orm:"column(name); size(32); default(); description(管理员姓名)"`
	AppIds string `orm:"column(app_ids); size(128); default(); description(拥有的应用 超管默认为空，拥有全部)"`
	Grade uint8 `orm:"column(grade); default(1); description(管理员等级 1:超管 2:普通)"`
	Staus uint8 `orm:"column(status); default(1); description(状态 1:正常 2:冻结)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime)"`
}

// 通过用户名获取用户信息
//func GetAdminInfoByName(userName string) {
//	admin := new(Admin)
//	a := orm.NewOrm().QueryTable(admin).Filter("name", userName)
//	beego.Info(a);
//}