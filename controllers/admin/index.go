package admin

import (
	"Go-Cms/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type resp struct {
	Stat uint8 `json:"stat"`
	Msg string `json:"msg"`
}

// 后台首页 登录页面
func (c *IndexController) Index() {
	c.Layout = "admin/index.html"
	c.TplName = "admin/header.html"
}

// 登录
func (c *IndexController) Login() {

	var stat uint8 = 1
	var msg string

	// 获取用户名
	userName := c.GetString("username", "")
	// 获取密码
	passWord := c.GetString("password", "")

	// 用户名不能为空
	if userName == "" {
		msg = "用户名不能为空"
		stat = 0
	}

	// 密码不能为空
	if len(passWord) != 6 {
		msg = "密码位数不合法"
		stat = 0
	}

	// 通过用户名获取用户信息
	adminData, err := models.GetAdminInfoByName(userName)

	// 用户不存在
	if err != nil {
		msg = "管理员不存在"
		stat = 0
	}

	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(passWord)))
	if md5Str != adminData.PassWord {
		msg = "用户名或者密码错误"
		stat = 0
	}
	msg = "登录成功"

	token := c.createToken(userName, passWord)
	// 设置cookie
	c.setLoginCookie(userName, token)
	// 设置session
	c.setLoginSession(userName, token)

	c.Data["json"] = resp{
		stat,
		msg,
	}
	c.ServeJSON()
}

// 退出
func (c *IndexController) Logout() {

}

// 生成token
func (c *IndexController) createToken(userName, passWord string) (token string){
	// 拼接密钥
	str := KEY + userName + passWord
	// md5 生成token
	token = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return
}

// 设置ssesion
func (c *IndexController) setLoginSession(userName, token string) {
	c.SetSession("userName", userName)
	c.SetSession("token", token)
}

// 设置cookie
func (c *IndexController) setLoginCookie(userName, token string) {
	c.Ctx.SetCookie("userName", userName, 3600, "/")
	c.Ctx.SetCookie("token", token, 3600, "/")
}
