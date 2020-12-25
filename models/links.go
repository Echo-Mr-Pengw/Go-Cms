// 友情链接

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Links struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Name       string    `orm:"column(name); size(32); default(); description(友情名);"form:"name"`
	Url        string    `orm:"column(url); size(64); default(); description(链接地址);"form:"url"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (u *Links) TableEngine() string {
	return "INNODB"
}

// 获取所有的友情链接信息
func GetAllLinks() (links []Links, num int64) {
	links = []Links{}
	num, _ = orm.NewOrm().QueryTable(new(Links)).OrderBy("-id").All(&links)
	return
}

//通过id,获取一条数据
func GetLinksInfoById(id string) (links *Links, err error) {
	links = new(Links)
	err = orm.NewOrm().QueryTable(links).Filter("id", id).One(links, "id", "name", "url", "status")
	return
}

// 更新友情链接信息
func UpdateLinksById(id, name, url, status string) (updateRows int64) {
	updateRows, err := orm.NewOrm().QueryTable(new(Links)).Filter("id", id).Update(orm.Params{
		"name": name,
		"url": url,
		"status": status,
	})

	if err != nil {
		updateRows = 0
	}
	return
}

// 添加友情链接
func AddLinks(links *Links) (int64) {
	addRow, err := orm.NewOrm().Insert(links)
	if err != nil {
		addRow = 0
	}
	return addRow
}