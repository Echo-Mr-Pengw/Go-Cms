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

    // 友情链接
    beego.Router("/links/list", &admin.LinksController{}, "get:List")
    beego.Router("/links/add", &admin.LinksController{}, "get,post:Add")
    beego.Router("/links/edit", &admin.LinksController{}, "get,post:Edit")

    // 个人简介
    beego.Router("/profile/list", &admin.ProfileController{}, "get:List")
    beego.Router("/profile/add", &admin.ProfileController{}, "get,post:Add")
    beego.Router("/profile/edit", &admin.ProfileController{}, "get,post:Edit")

    // 标签列表
    beego.Router("tag/list", &admin.TagController{}, "get:List")
    beego.Router("tag/add", &admin.TagController{}, "get,post:Add")
    beego.Router("tag/edit", &admin.TagController{}, "get,post:Edit")

    // 文章管理
    beego.Router("article/list", &admin.ArticleController{}, "get:List")
    beego.Router("article/add", &admin.ArticleController{}, "get,post:Add")
    beego.Router("article/edit", &admin.ArticleController{}, "get,post:Edit")
}
