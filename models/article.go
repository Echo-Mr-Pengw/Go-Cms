// 文章

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	TagId      uint      `orm:"column(tag_id); default(0); description(标签id);"form:"tag_id"`
	Title      string    `orm:"column(title); size(648); default(); description(文章标题);"form:"title"`
	Abstract   string    `orm:"column(abstract); size(128); default(); description(文章摘要);"form:"abstract"`
	Content    string    `orm:"column(content); type(text); default(); description(文章内容);"form:"content"`
	Author     string    `orm:"column(author); size(32); default(); description(文章作者);"form:"author"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (u *Article) TableEngine() string {
	return "INNODB"
}

// 获取所有的文章
func GetArticleList() (art []Article, num int64) {
	art = []Article{}
	num, _ = orm.NewOrm().QueryTable(new(Article)).OrderBy("-id").All(&art)
	return
}

// 获取正常的文章
//func GetNormalArticleList() (art []Article, num int64) {
//	art = []Article{}
//	num, _ = orm.NewOrm().QueryTable(new(Article)).Filter("status", 1).OrderBy("-id").All(&art)
//	return
//}

// 连表查询正常显示的文章
func GetNormalArticleList(startArticleId, endArticleId int) ([]orm.Params) {
	var list []orm.Params
	sql := "SELECT  a.id,a.title,a.abstract,a.author,a.create_time, ifnull(b.read_num, 0) as read_num FROM tbl_article as a LEFT JOIN tbl_article_read_num as b ON a.id = b.article_id WHERE a.id>? AND a.id<=?"
	orm.NewOrm().Raw(sql, startArticleId, endArticleId).Values(&list)
	return list
}

// 获取正常文章显示文章的总数量
func GetNormalArticleNum() (int64){
	total, _ := orm.NewOrm().QueryTable(new(Article)).Filter("status", 1).Count()
	return total
}

// 添加文章
func AddArticle(a *Article) (addRow int64) {
	addRow, err := orm.NewOrm().Insert(a)
	if err != nil {
		addRow = 0
	}
	return
}

// 通过文章id,获取数据
func GetArtcileInfoById (id string) (art *Article, err error) {
	art = new(Article)
	err = orm.NewOrm().QueryTable(art).Filter("id", id).One(art, "id", "tag_id", "title", "content", "author", "status")
	return
}

// 更新文章信息
func UpdateArticleById(id, tag_id, title, content, author, status string) (updateRows int64) {
	updateRows, err := orm.NewOrm().QueryTable(new(Article)).Filter("id", id).Update(orm.Params{
		"tag_id": tag_id,
		"title": title,
		"content": content,
		"author": author,
		"status": status,
	})

	if err != nil {
		updateRows = 0
	}
	return
}

// 最近10篇新文章
func GetTopTenNormalArticle() (art []Article, num int64) {
	art = []Article{}
	num, _ = orm.NewOrm().QueryTable(new(Article)).Filter("status", 1).OrderBy("-id").Limit(10).All(&art, "id", "title")
	return
}

// 通过文章id获取点击量前10的文章
func GetTopTenNormalArticleByArtId(articleId []uint) (art []Article, num int64) {
	art = []Article{}
	num, _ = orm.NewOrm().QueryTable(new(Article)).Filter("id__in", articleId).All(&art, "id", "title")
	return
}
