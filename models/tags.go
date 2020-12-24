// 标签

package models

import (
	"time"
)

type ArticleTag struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Name       string     `orm:"column(name); size(32); default(); description(标签名);"form:"name"`
	Total      uint       `orm:"column(total); default(0); description(标签下文章的总数);"form:"total"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (u *ArticleTag) TableEngine() string {
	return "INNODB"
}