package home

import (
	"Go-Cms/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {

	// 个人简介
	profile,_ := models.GetOneProfile()
	// 友情正常显示的连接
	links, _ := models.GetNormalLinks()
	// 随笔分类
	tag, _ := models.GetNormalTagsList()
	// 文章列表
	articleList,_ := models.GetArticleList()

	c.Data["links"]       = links
	c.Data["tag"]         = tag
	c.Data["profile"]     = profile
	c.Data["articleList"] = articleList
	c.TplName = "home/index.html"
}

