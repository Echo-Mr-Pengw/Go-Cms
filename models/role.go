// 角色

package models

import "time"

type Role struct {
	Id uint `orm:"column(id); pk; auto; description(主键)"`
	ResourceIds string `orm:"column(resource_ids); size(1024); default(); description(资源id集)"`
	Name string `orm:"column(name); size(128); default(); description(角色名)"`
	Status uint8 `orm:"column(status); size(1); default(1); description(角色状态 1:正常 2:冻结"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *Role) TableEngine() string {
	return "INNODB"
}