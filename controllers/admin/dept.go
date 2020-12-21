package admin

import (
	"Permission-Platform/models"
)

type DeptController struct {
	BaseController
}

func (c *DeptController) List() {
	deptList,_ := models.GetAllDeptList()

	c.Data["statusMap"] = statusMap
	c.Data["deptList"]  = deptList

	c.TplName = "dept/list.html"
}
