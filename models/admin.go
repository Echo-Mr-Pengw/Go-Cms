package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Admin struct {
	Id uint `orm:"column(id); pk; auto; description(主键)"`
	Emplid string `orm:"column(emplid); size(32); default(); description(管理员工号)"`
	Name string `orm:"column(name); size(32); default(); description(管理员姓名)"`
	AppIds string `orm:"column(app_ids); size(128); default(); description(拥有的应用 超管默认为空，拥有全部)"`
	PassWord string `orm:"column(password); size(32); deafult(); description(密码)"`
	Grade uint8 `orm:"column(grade); default(1); description(管理员等级 1:超管 2:普通)"`
	Staus uint8 `orm:"column(status); default(1); description(状态 1:正常 2:冻结)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime)"`
}

// 定义表的存储引擎
func (u *Admin) TableEngine() string {
	return "INNODB"
}

// 通过用户名获取用户信息
func GetAdminInfoByName(userName string) (*Admin, error) {
	admin := new(Admin)
	err := orm.NewOrm().QueryTable(admin).Filter("name", userName).Filter("status", 1).One(admin, "password")
	return admin, err
}