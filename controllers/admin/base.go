package admin

import "github.com/astaxie/beego"

// '继承'
type BaseController struct {
	beego.Controller
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