// 用户

package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id uint `json:"id",orm:"column(id); pk; auto; description(主键 用户id)"`
	Deptid uint `json:"deptid",orm:"column(dept_id); description(部门id)"`
	Name string `json:"name",orm:"column(name); size(64); default(); description(用户名)"`
	Status uint8 `json:"status",orm:"column(status); size(1); default(1); description(状态)"`
	CreateTime time.Time `json:"create_time",orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *User) TableEngine() string {
	return "INNODB"
}

// 通过用户名获取有效的用户（模糊匹配）
func GetNormalUserByName(userName string) ([]User, int64) {
	 var user []User
	 num, _ := orm.NewOrm().QueryTable(new(User)).Filter("status", 1).Filter("name__istartswith", userName).All(&user, "id", "name")
	 return user, num
}