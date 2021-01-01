package admin

import "Go-Cms/models"

type LinksController struct {
	BaseController
}

func (c *LinksController) List() {

	list, _ := models.GetAllLinks()

	c.LayoutSections = make(map[string]string)

	c.Data["linksList"] = list
	c.Data["statusMap"] = statusMap

	c.Layout  = "admin/links/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"]    = "admin/nav.html"
	c.LayoutSections["footer"] = "admin/footer.html"
}

func (c *LinksController) Add() {

	if c.IsPost() {
		links := models.Links{}
		if err := c.ParseForm(&links); err != nil {
			c.ResponseJson(0, "参数解析错误", "")
		}

		addRow := models.AddLinks(&links)
		if addRow == 0 {
			c.ResponseJson(0, "添加失败", "")
		}
		c.ResponseJson(1, "添加成功", "")
	}

	c.Data["statusMap"] = statusMap
	c.Layout   = "admin/links/add.html"
	c.TplName  = "admin/header.html"
}

func (c *LinksController) Edit() {

	if c.IsPost() {
		id     := c.GetString("id")
		name   := c.GetString("name", "")
		url    := c.GetString("url", "")
		status := c.GetString("status", "1")

		updateRow := models.UpdateLinksById(id, name, url, status)
		if updateRow == 0 {
			c.ResponseJson(0, "更新失败", "")
		}
		c.ResponseJson(1, "更新成功", "")
	}

	id := c.GetString("id", "1")
	info, _ := models.GetLinksInfoById(id)

	c.Data["info"] = info
	c.Data["statusMap"] = statusMap
	c.Layout   = "admin/links/edit.html"
	c.TplName  = "admin/header.html"
}
