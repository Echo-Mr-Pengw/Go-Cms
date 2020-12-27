package admin

import (
	"github.com/astaxie/beego"
	"strings"
)

// 继承
type BaseController struct {
	beego.Controller
}

// 定义状态常量及密钥
const (
	NORMAL = 1
	FROZEN = 2
	KEY    = "482d909db559cd0922da137487d6f7e4"
)

// 状态定义
var statusMap = map[uint8]string{
	1 : "正常",
	2 : "冻结",
}

// 管理员等级
var adminLevel = map[uint8]string {
	1: "超级管理员",
	2: "普通管理员",
}

// 图片的格式
var imgExtName = map[string]string {
	".jpg": ".jpg",
	".jpeg": ".jepg",
	".png": ".png",
}

// 定义返回消息体
type Response struct {
	Stat uint8 `json:"stat"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

// 检查session
func (c *BaseController) Prepare() {
	cookieToken    := c.Ctx.GetCookie("token")
	cookieUserName := c.Ctx.GetCookie("userName")
	if cookieUserName == "" || cookieToken == "" {
		c.Redirect("/", 302)
	}

	sessToken    := c.GetSession("token")
	sessUserName := c.GetSession("userName")
	if sessUserName == "" || sessToken == "" {
		c.Redirect("/", 302)
	}

	if (cookieUserName != sessUserName) || (cookieToken != sessToken) {
		c.Redirect("/", 302)
	}
}

// 定义返回的json
func (c *BaseController) ResponseJson(stat uint8, msg string, data interface{}) {
	resp := Response{
		Stat: stat,
		Msg:  msg,
		Data: data,
	}

	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) IsPost() bool {
	if strings.ToLower(c.Ctx.Request.Method) == "post" {
		return true
	}
	return false
}