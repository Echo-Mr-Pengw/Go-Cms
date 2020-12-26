package admin

import (
	"Go-Cms/models"
	"math/rand"
	"path"
	"strconv"
	"time"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) List() {

	profileList, num := models.GetProfileList()

	c.Data["num"]         = num
	c.Data["statusMap"]   = statusMap
	c.Data["profileList"] = profileList

	c.TplName = "admin/profile/list.html"
}

func (c *ProfileController) Add() {

	if c.IsPost() {

		f, h, err := c.GetFile("file")
		defer f.Close()

		if err != nil {
			c.ResponseJson(0, "图片上传失败", "")
		}

		// 获取图片的扩展名
		extName := path.Ext(h.Filename)
		if _, isOk := imgExtName[extName]; !isOk {
			c.ResponseJson(0, "文件格式必须为：jpg、jpeg、png", "")
		}

		//dir, _ := os.Getwd()

		rand.Seed(time.Now().UnixNano())
		rand := rand.Intn(1000)

		imgSavePath :=  "/static/images/wx_qq/" + strconv.Itoa(rand) + extName
		err = c.SaveToFile("file", imgSavePath)

		if err != nil {
			c.ResponseJson(0, "图片上传失败", "")
		}

		profile := models.Profile{}
		if err := c.ParseForm(&profile); err != nil {
			c.ResponseJson(0, "参数解析失败", "")
		}

		profile.WxQqImg = imgSavePath

		if models.AddProfile(&profile) == 0 {
			c.ResponseJson(0, "添加失败", "")
		}
		c.ResponseJson(1, "添加成功", "")
	}

	c.Data["statusMap"]   = statusMap
	c.TplName = "admin/profile/add.html"
}

func (c *ProfileController) Edit() {

}