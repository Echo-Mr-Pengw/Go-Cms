// 友情链接

package models

import (
	"time"
)

type Links struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Name       string    `orm:"column(name); size(32); default(); description(友情名);"form:"name"`
	Url        string    `orm:"column(url); size(64); default(); description(链接地址);"form:"motto"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (u *Links) TableEngine() string {
	return "INNODB"
}