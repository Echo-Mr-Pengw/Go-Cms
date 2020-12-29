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
	// 文章列表
	articleList,_ := models.GetArticleList()

	c.Data["profile"]     = profile
	c.Data["articleList"] = articleList
	c.TplName = "home/index.html"
}

