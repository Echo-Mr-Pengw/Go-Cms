package admin

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.TplName = "admin/main.html"
}