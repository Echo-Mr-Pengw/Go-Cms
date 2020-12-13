package admin

import (
	"Permission-Platform/models"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) List() {
	// 管理员列表
	adminList, _ := models.GetAdminList()
	c.Data["statusMap"]     = statusMap
	c.Data["adminLevelMap"] = adminLevel
	c.Data["adminList"]     = adminList

	c.TplName = "admin/list.html"
}

func (c *AdminController) Add() {

	if c.IsPost() {

	}

	appList, _ := models.GetAllAppList()

	c.Data["appList"]       = appList
	c.Data["statusMap"]     = statusMap
	c.Data["adminLevelMap"] = adminLevel
	c.TplName = "admin/add.html"
}

func (c *AdminController) Edit() {
	c.TplName = "admin/edit.html"
}