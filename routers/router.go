package routers

import (
	"Permission-Platform/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	// 后端首页
    beego.Router("/", &admin.IndexController{}, "get:Index")
    // 登录操作
    beego.Router("/login", &admin.IndexController{}, "post:Login")
    // 退出操作
    beego.Router("/logout", &admin.IndexController{}, "get:Logout")

    // 登录成功主页
    beego.Router("/main", &admin.MainController{}, "get:Index")
}
