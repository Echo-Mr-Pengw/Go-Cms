package admin

import (
	"Go-Cms/models"
	"crypto/md5"
	"fmt"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) List() {
	// 管理员列表
	adminList, _ := models.GetAdminList()

	c.Data["statusMap"]      = statusMap
	c.Data["adminLevelMap"]  = adminLevel
	c.Data["adminList"]      = adminList

	c.Layout  = "admin/admin/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"]    = "admin/nav.html"
	c.LayoutSections["footer"] = "admin/footer.html"
}

func (c *AdminController) Add() {

	if c.IsPost() {
		admin := models.Admin{}
		if err := c.ParseForm(&admin); err != nil {
			c.ResponseJson(0, "解析失败", "")
		}
		admin.PassWord  = fmt.Sprintf("%x", md5.Sum([]byte(admin.PassWord)))
		insertRow := models.AddAdminInfo(&admin)
		if insertRow < 0 {
			c.ResponseJson(0, "添加失败", "")
		}
		c.ResponseJson(1, "添加成功", "")
	}

	c.LayoutSections = make(map[string]string)

	c.Data["statusMap"]      = statusMap
	c.Data["adminLevelMap"]  = adminLevel

	c.Layout   = "admin/admin/add.html"
	c.TplName  = "admin/header.html"
}

func (c *AdminController) Edit() {

	if c.IsPost() {
		id         := c.GetString("id", "")
		grade      := c.GetString("grade", "", "")
		status     := c.GetString("status", "")

		// 参数验证
		if id == "" || grade == "" || status == "" {
			c.ResponseJson(0, "参数错误", "")
		}

		updateRows := models.UpdateAdminInfoById(id, grade, status)
		if updateRows == 0 {
			c.ResponseJson(0, "编辑失败", "")
		}
		c.ResponseJson(1, "编辑成功", "")
	}

	id      := c.GetString("id", "")
	info,_  := models.GetAdminInfoById(id)

	c.Data["info"]          = info
	c.Data["statusMap"]     = statusMap
	c.Data["adminLevelMap"] = adminLevel

	c.Layout   = "admin/admin/edit.html"
	c.TplName  = "admin/header.html"
}