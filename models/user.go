// 用户

package models

import "time"

type User struct {
	Id uint `orm:"column(id); pk; auto; description(主键 用户id)"`
	Deptid uint `orm:"column(dept_id); description(部门id)"`
	Name string `orm:"column(name); size(64); default(); description(用户名)"`
	Status uint8 `orm:"column(status); size(1); default(1); description(状态)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *User) TableEngine() string {
	return "INNODB"
}