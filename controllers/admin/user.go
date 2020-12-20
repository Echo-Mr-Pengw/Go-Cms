package admin

import (
	"Permission-Platform/models"
	"fmt"
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

// 用户列表
func (c *UserController) List () {
	userList,_ := models.GetAllUserList()
	fmt.Println(userList)
	c.Data["statusMap"] = statusMap
	c.Data["userList"] = userList
	c.TplName = "user/list.html"
}