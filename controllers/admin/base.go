package admin

import (
	"github.com/astaxie/beego"
	"strings"
)

// '继承'
type BaseController struct {
	beego.Controller
}

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

// 定义返回消息体
type Response struct {
	Stat uint8 `json:"stat"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
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
}

func (c *BaseController) IsPost() bool {
	if strings.ToLower(c.Ctx.Request.Method) == "post" {
		return true
	}
	return false
}