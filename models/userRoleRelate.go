// 用户与角色之间的关系

package models

import "time"

type UserRoleRelate struct {
	Id uint `orm:"column(id); pk; auto; description(主键)"`
	AppId uint `orm:"column(app_id); default(0); description(应用id)"`
	Emplid uint `orm:"column(emplid); default(0); description(工号)"`
	RoleId uint `orm:"column(role_id); default(0); description(角色id)"`
	Status uint8 `orm:"column(status); default(1); description(员工所属角色的状态 1:正常 2:失效)"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间)"`
}