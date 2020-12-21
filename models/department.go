// 部门

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

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

// 获取所有的部门
func GetAllDeptList() ([]Department, int64){
	var dept []Department
	num, _ := orm.NewOrm().QueryTable(new(Department)).OrderBy("-id").All(&dept)
	return dept, num
}