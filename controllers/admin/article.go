package admin

import (
	"Go-Cms/models"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {

	// 文章标签map
	var tagMap = make(map[uint]string)

	tagList, _ := models.GetTagsList()
	for _, v := range tagList {
		tagMap[v.Id] = v.Name
	}


	articleList, _ := models.GetArticleList()

	c.Data["tagMap"]      = tagMap
	c.Data["articleList"] = articleList
	c.Data["statusMap"]   = statusMap
	c.TplName = "admin/article/list.html"
}

func (c *ArticleController) Add() {

	if c.IsPost() {
		article := models.Article{}
		if err := c.ParseForm(&article); err != nil {
			c.ResponseJson(0, "参数解析失败", "")
		}

		if models.AddArticle(&article) == 0 {
			c.ResponseJson(0, "添加失败", "")
		}
		c.ResponseJson(1, "添加成功", "")
	}

	// 获取有效的文章标签
	tagList, _ := models.GetNormalTagsList()

	c.Data["tagList"]   = tagList
	c.Data["statusMap"] = statusMap
	c.TplName = "admin/article/add.html"
}

func (c *ArticleController) Edit() {

	id := c.GetString("id")

	if c.IsPost() {
		id      := c.GetString("id")
		title   := c.GetString("title")
		content := c.GetString("content")
		author  := c.GetString("author")
		status  := c.GetString("status")
		tag_id  := c.GetString("tag_id")

		if models.UpdateArticleById(id, tag_id, title, content, author, status) == 0 {
			c.ResponseJson(0, "更新失败", "")
		}
		c.ResponseJson(1, "更新成功", "")
	}

	// 通过文章id,获取文章信息
	artList, _ := models.GetArtcileInfoById(id)

	// 获取有效的文章标签
	tagList, _ := models.GetNormalTagsList()

	c.Data["tagList"]   = tagList
	c.Data["artList"]   = artList
	c.Data["statusMap"] = statusMap
	c.TplName = "admin/article/edit.html"
}