package admin

import (
	"Go-Cms/models"
	"crypto/md5"
	"fmt"
)

type IndexController struct {
	BaseController
}

// 后台首页 登录页面
func (c *IndexController) Index() {
	c.TplName = "admin/index.html"
}

// 登录
func (c *IndexController) Login() {

	// 获取用户名
	userName := c.GetString("username", "")
	// 获取密码
	passWord := c.GetString("password", "")

	// 用户名不能为空
	if userName == "" {
		c.ResponseJson(0, "用户名不能为空", "")
	}

	// 密码不能为空
	if len(passWord) != 6 {
		c.ResponseJson(0, "密码位数不合法", "")
	}

	// 通过用户名获取用户信息
	adminData, err := models.GetAdminInfoByName(userName)

	// 用户不存在
	if err != nil {
		c.ResponseJson(0, "管理员不存在", "")
	}

	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(passWord)))
	if md5Str != adminData.PassWord {
		c.ResponseJson(0, "用户名或者密码错误", "")
	}

	c.ResponseJson(1, "登录成功", "")
}

// 退出
func (c *IndexController) Logout() {

}