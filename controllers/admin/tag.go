package admin

import "Go-Cms/models"

type TagController struct {
	BaseController
}

func (c *TagController) List() {

	tagList, _ := models.GetTagsList()
	c.Data["tagList"]   = tagList
	c.Data["statusMap"] = statusMap
	c.TplName = "admin/tag/list.html"
}

func (c *TagController) Add() {

	if c.IsPost() {

		tag := models.ArticleTag{}
		if err := c.ParseForm(&tag); err != nil {
			c.ResponseJson(0, "参数解析失败", "")
		}

		if models.AddTag(&tag) == 0 {
			c.ResponseJson(0, "添加失败", "")
		}
		c.ResponseJson(1, "添加成功", "")
	}

	c.Data["statusMap"] = statusMap
	c.TplName = "admin/tag/add.html"
}

func (c *TagController) Edit() {

	id := c.GetString("id")
	tagInfo, _ := models.GetTagInfoById(id)

	if c.IsPost() {
		id     := c.GetString("id")
		name   := c.GetString("name")
		status := c.GetString("status")

		row := models.UpdateTagById(id, name, status)
		if row == 0 {
			c.ResponseJson(0, "编辑失败", "")
		}
		c.ResponseJson(1, "编辑成功", "")
	}

	c.Data["tagInfo"]   = tagInfo
	c.Data["statusMap"] = statusMap
	c.TplName = "admin/tag/edit.html"

}