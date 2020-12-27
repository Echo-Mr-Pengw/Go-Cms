// 标签

package models

import (
	"github.com/astaxie/beego/orm"
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
func (a *ArticleTag) TableEngine() string {
	return "INNODB"
}

// 获取所有标签
func GetTagsList() (tag []ArticleTag, num int64) {
	tag = []ArticleTag{}
	num, _ = orm.NewOrm().QueryTable(new(ArticleTag)).OrderBy("-id").All(&tag)
	return
}

// 获取有效的标签
func GetNormalTagsList() (tag []ArticleTag, num int64) {
	tag = []ArticleTag{}
	num, _ = orm.NewOrm().QueryTable(new(ArticleTag)).Filter("status", 1).OrderBy("-id").All(&tag)
	return
}

func AddTag(a *ArticleTag) (addRow int64) {
	addRow, err := orm.NewOrm().Insert(a)
	if err != nil {
		addRow = 0
	}
	return
}

// 获取一条记录
func GetTagInfoById(id string) (tag *ArticleTag, err error) {
	tag = new(ArticleTag)
	err = orm.NewOrm().QueryTable(tag).Filter("id", id).One(tag, "id", "name", "status")
	return
}

// 更新标签信息
func UpdateTagById(id, name, status string) (updateRows int64) {
	updateRows, err := orm.NewOrm().QueryTable(new(ArticleTag)).Filter("id", id).Update(orm.Params{
		"name": name,
		"status": status,
	})

	if err != nil {
		updateRows = 0
	}
	return
}