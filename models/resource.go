// 资源

package models

import "time"

type Resource struct {
	Id uint `orm:"column(id); pk; auto; description(主键)"`
	Appid uint `orm:"column(app_id); default(0); description(应用id)"`
	ParentId uint `orm:"column(parent_id); default(0); description(父级id)"`
	Path string `orm:"column(path); size(128); default(); description(父级路径)"`
	Name string `orm:"column(name); size(128); default(); description(资源名)"`
	Status uint8 `orm:"column(status); size(1); default(1); description(状态: 1:正常 2:冻结)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *Resource) TableEngine() string {
	return "INNODB"
}