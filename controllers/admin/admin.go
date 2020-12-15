package admin

import (
	"Permission-Platform/models"
	"crypto/md5"
	"fmt"
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

	appList, _ := models.GetAllAppList()

	c.Data["appList"]       = appList
	c.Data["statusMap"]     = statusMap
	c.Data["adminLevelMap"] = adminLevel
	c.TplName = "admin/add.html"
}

func (c *AdminController) Edit() {
	c.TplName = "admin/edit.html"
}