// 管理员

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Admin struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Name       string    `orm:"column(name); size(32); default(); description(管理员姓名);"form:"name"`
	PassWord   string    `orm:"column(password); size(32); deafult(); description(密码);"form:"password"`
	Grade      uint8     `orm:"column(grade); default(1); description(管理员等级 1:超管 2:普通);"form:"grade"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
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

// 获取所有的管理员数据
func GetAdminList() (admin []Admin, num int64) {
	admin = []Admin{}
	num, _ = orm.NewOrm().QueryTable(new(Admin)).OrderBy("-id").All(&admin)
	return
}

// 添加管理员
func AddAdminInfo(app *Admin) (id int64) {
	id, _ = orm.NewOrm().Insert(app)
	return
}

// 通过该管理员id,获取管理员信息
func GetAdminInfoById(id string) (*Admin, error){
	admin := new(Admin)
	err := orm.NewOrm().QueryTable(admin).Filter("id", id).One(admin, "id", "name", "grade", "status")
	return admin, err
}

// 通过管理员id,更新管理员信息
func UpdateAdminInfoById(id, grade, status string) int64 {
	updateRows, err := orm.NewOrm().QueryTable(new(Admin)).Filter("id", id).Update(orm.Params{
		"grade": grade,
		"status": status,
	})

	if err != nil {
		updateRows = 0
	}
	return updateRows
}