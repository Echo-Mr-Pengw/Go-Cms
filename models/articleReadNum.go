// 文章阅读量
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type ArticleReadNum struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"`
	ArticleId  uint      `orm:"column(article_id); unique; description(文章id);"`
	ReadNum    uint      `orm:"column(read_num); default(0); description(阅读量);"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"`
}

// 定义表的存储引擎
func (u *ArticleReadNum) TableEngine() string {
	return "INNODB"
}

// 新增/更新浏览数
func UpdateReadNum(articleId string) {
	sql := "insert into tbl_article_read_num(article_id,read_num) values(?,1) on duplicate key update read_num=read_num+1"
	_, err := orm.NewOrm().Raw(sql, articleId).Exec()
	fmt.Println(111, err)
}

// 获取前10阅读数最高的文章id
func GetTopTenArtIdByReadNum() (readNum []ArticleReadNum, num int64){
	readNum = []ArticleReadNum{}
	num, _ = orm.NewOrm().QueryTable(new(ArticleReadNum)).Filter("status", 1).OrderBy("-read_num").Limit(10).All(&readNum, "article_id", "read_num")
	return
}