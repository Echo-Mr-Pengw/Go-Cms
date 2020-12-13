package admin

import (
	"Permission-Platform/models"
	"strconv"
	"strings"
)
type AppController struct {
	BaseController
}

// 获取全部的应用
func (c *AppController) GetAppList() {
	appList, _ := models.GetAllAppList()
	c.Data["appList"]   = appList
	c.Data["statusMap"] = statusMap
	c.TplName = "app/list.html"
}

// 添加应用
func (c *AppController) Add() {

	if strings.ToLower(c.Ctx.Request.Method) == "post" {
		app := models.App{}
		if err := c.ParseForm(&app); err != nil {
			c.ResponseJson(0, "参数解析失败", "")
		}

		addRow := models.AddAppInfo(&app)
		if addRow == 0 {
			c.ResponseJson(0 , "添加失败", "")
		}
		c.ResponseJson(1, "添加成功", "")
	}

	c.Data["statusMap"] = statusMap
	c.TplName = "app/add.html"
}

// 编辑应用
func (c *AppController) Edit() {

	// 获取当前url中的id
	id, _ := strconv.ParseUint(c.GetString("id"), 10, 0)
	app   := models.App{}
	app.Id = uint(id)

	if strings.ToLower(c.Ctx.Request.Method) == "post" {
		id, _ := strconv.ParseUint(c.GetString("id"), 10, 0)
		app := models.App{}
		app.Id = uint(id)

		if err := models.GetDetailAppData(&app); err != nil {
			c.ResponseJson(0, "数据获取失败", "")
		}

		app.Name   = c.GetString("name")
		status, _ := strconv.ParseUint(c.GetString("status"), 10, 0)
		app.Status = uint8(status)

		if err := models.UpdateAppData(&app); err != nil {
			c.ResponseJson(0, "更新失败", "")
		}
		c.ResponseJson(1, "更新成功", "")
	}
	// 通过主键获取对应的数据
	models.GetDetailAppData(&app)

	c.Data["detailData"] = app
	c.Data["statusMap"]  = statusMap
	c.TplName = "app/edit.html"
}