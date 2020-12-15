package admin

import (
	"Permission-Platform/models"
)

type UserController struct {
	BaseController
}

// 通过用户名进行模糊搜索
func (c *UserController) SearchNormalUser() {

	name := c.GetString("name", "")
	if name == "" {
		c.ResponseJson(0, "管理员不能为空", "")
	}

	list, _ := models.GetNormalUserByName(name)
	c.ResponseJson(1, "sucess", list)
}