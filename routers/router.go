package routers

import (
	"Go-Cms/controllers/admin"
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

    // 管理员相关路由
    // 管理员列表
    beego.Router("/admin/list", &admin.AdminController{}, "get,post:List")
    // 管理员添加
    beego.Router("/admin/add", &admin.AdminController{}, "get,post:Add")
    // 管理员编辑
    beego.Router("/admin/edit", &admin.AdminController{}, "get,post:Edit")
}
