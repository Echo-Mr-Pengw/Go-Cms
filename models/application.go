// 权限申请单

package models

import "time"

type Application struct {
	Id uint `orm:"column(id); pk; auto; description(主键)"`
	Appid uint `orm:"column(app_id); default(0); description(应用id)"`
	AppName string `orm:"column(app_name); size(128); default(); description(应用名)"`
	Emplid uint `orm:"column(emplid); default(0); description(申请人的工号)"`
	Deptid uint `orm:"column(deptid); default(0); description(申请人的部门)"`
	AppliedEmplid string `orm:"column(applied_emplid); size(648); default(); description(被申请人的工号)"`
	AppliedNames string `orm:"column(applied_names); size(648); default(); description(被申请人/部门名)"`
	RoleIds string `orm:"column(role_ids); size(648); default(); description(角色集合)"`
	ApplyType uint8 `orm:"column(apply_type); size(1); default(1); description(申请类型 1:用户 2:部门)"`
	ApplyStatus uint8 `orm:"column(apply_status); size(1); default(1); description(申请状态 1:待提审 2:审批中 3:撤回 4:通过 5:驳回)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}

// 定义表的存储引擎
func (u *Application) TableEngine() string {
	return "INNODB"
}