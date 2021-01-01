package admin

import (
	"Go-Cms/models"
	"fmt"
	"strconv"
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

	c.Layout  = "admin/article/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"]    = "admin/nav.html"
	c.LayoutSections["footer"] = "admin/footer.html"
}

func (c *ArticleController) Add() {

	if c.IsPost() {
		article := models.Article{}
		if err := c.ParseForm(&article); err != nil {
			c.ResponseJson(0, "参数解析失败", "")
		}
		fmt.Println(article)
		article.Abstract = article.Content[0:435]

		if models.AddArticle(&article) == 0 {
			c.ResponseJson(0, "添加失败", "")
		}

		// 文章状态为正常：更新该标签下的文章数量
		switch article.Status {
			// 正常
			case NORMAL:
				models.AddArticleNumByTagId(article.TagId)
		}

		c.ResponseJson(1, "添加成功", "")
	}

	// 获取有效的文章标签
	tagList, _ := models.GetNormalTagsList()

	c.Data["tagList"]   = tagList
	c.Data["statusMap"] = statusMap

	c.Layout  = "admin/article/add.html"
	c.TplName = "admin/header.html"
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

		// 文章状态为正常：更新该标签下的文章数量
		intStatus, _ := strconv.Atoi(status)
		intTagId, _ := strconv.ParseUint(tag_id, 10, 0)
		switch intStatus {
		// 正常
		case NORMAL:
			models.AddArticleNumByTagId(uint(intTagId))
		// 冻结
		case FROZEN:
			models.MinusArticleNumByTagId(uint(intTagId))
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

	c.Layout  = "admin/article/edit.html"
	c.TplName = "admin/header.html"
}