// 部门

package models

import "time"

type Department struct {
	Id uint `orm:"column(id); pk; auto; description(主键 部门id)"`
	Name string `orm:"column(name); size(128); default(); description(部门名)"`
	Status uint8 `orm:"column(status); size(1); default(1); description(状态: 1:正常 2:冻结)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *Department) TableEngine() string {
	return "INNODB"
}