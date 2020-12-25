// 个人简介

package models

import (
	"time"
)

type Profile struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Nickname   string    `orm:"column(nickname); size(32); default(); description(昵称);"form:"nickname"`
	Motto      string    `orm:"column(motto); size(64); default(); description(座右铭);"form:"motto"`
	WxQqImg    string    `orm:"column(wx_qq_img); size(64); default(); description(个人微信/qq图片);"form:"wx_qq_img"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (p *Profile) TableEngine() string {
	return "INNODB"
}