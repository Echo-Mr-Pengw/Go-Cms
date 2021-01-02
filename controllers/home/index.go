package home

import (
	"Go-Cms/models"
	"github.com/astaxie/beego"
	"strconv"
)

const (
	pageNum = 10   // 每页的页数
)

type IndexController struct {
	beego.Controller
}

// 阅读量排名前10的文章数据
type topTenArticle struct {
	Id uint
	ReadNum uint
	Title string
}

func (c *IndexController) Index() {

	startPage := 0
	endPage   := 10

	c.Data["total"]         = c.getTotalArticleNum()     // 文章总数
	c.Data["links"]         = c.getLinks()               // 正常状态的友情链接
	c.Data["topTenArt"]     = c.getTopTenArt()           // 最近发布的前10的文章
	c.Data["tag"]           = c.getTagList()             // 标签列表
	c.Data["profile"]       = c.getProfile()             // 获取个人简介
	c.Data["articleList"]   = c.getArtList(startPage, endPage)             // 文章列表
	c.Data["topTenViewArt"] = c.getTopTenViewsArticle()  // 阅读量排名前10的文章


	c.Layout  = "home/index.html"
	c.TplName = "home/list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header"] = "home/header.html"
	c.LayoutSections["footer"] = "home/footer.html"
	c.LayoutSections["left"]   = "home/left.html"
}

// 正常显示的文章总数
func (c *IndexController) getTotalArticleNum() (int64){
	total := models.GetNormalArticleNum()
	return total
}

// 获取个人简介
func (c *IndexController) getProfile() (profile models.Profile) {
	// 个人简介
	profile,_ = models.GetOneProfile()
	return
}

// 正常状态的友情链接
func (c *IndexController) getLinks() (links []models.Links) {
	links, _ = models.GetNormalLinks()
	return
}

// 最近编写的前10的最新随笔
func (c *IndexController) getTopTenArt() (topTenArt []models.Article) {
	topTenArt, _ = models.GetTopTenNormalArticle()
	return
}

// 随笔分类
func (c *IndexController) getTagList() (tag []models.ArticleTag) {
	tag, _ = models.GetNormalTagsList()
	return
}

// 文章列表
func (c *IndexController) getArtList(startPage, endPage int) (articleList interface{}) {
	articleList = models.GetNormalArticleList(startPage, endPage)
	return
}

func (c *IndexController) getOneArticleById (id string) (art *models.Article){
	art, _ = models.GetArtcileInfoById(id)
	return
}

// 浏览量前10的文章
func (c *IndexController) getTopTenViewsArticle() ([]topTenArticle) {
	var articleId []uint
	var articleListSlice []topTenArticle
	var readNumMap = make(map[uint]uint, 10)

	readNum, num := models.GetTopTenArtIdByReadNum()
	if num == 0 {
		return articleListSlice
	}

	for _, v := range readNum {
		// 文章id保存到切片中
		articleId = append(articleId, v.ArticleId)
		// 文章id以及数量保存到map中
		readNumMap[v.ArticleId] = v.ReadNum
	}

	articleList, num := models.GetTopTenNormalArticleByArtId(articleId)
	if num == 0 {
		return articleListSlice
	}

	for _, v := range articleList {
		if readNum, ok := readNumMap[v.Id]; ok {
			data := topTenArticle {
				Id: v.Id,
				ReadNum: readNum,
				Title: v.Title,
			}
			articleListSlice = append(articleListSlice, data)
		}
	}
	return articleListSlice
}


func (c *IndexController) Detail() {
	id := c.Ctx.Input.Param(":id")

	c.Data["links"]          = c.getLinks()               // 正常状态的友情链接
	c.Data["topTenArt"]      = c.getTopTenArt()           // 最近发布的前10的文章
	c.Data["tag"]            = c.getTagList()             // 标签列表
	c.Data["profile"]        = c.getProfile()             // 获取个人简介
	//c.Data["articleList"]    = c.getArtList()             // 正常显示的文章数量
	c.Data["articleContent"] = c.getOneArticleById(id)    // 当前文章id的内容
	c.Data["topTenViewArt"]  = c.getTopTenViewsArticle()  // 阅读量排名前10的文章

	// 更新当前文章的阅读量
	models.UpdateReadNum(id)

	c.Layout  = "home/detail.html"
	c.TplName = "home/content.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header"] = "home/header.html"
	c.LayoutSections["footer"] = "home/footer.html"
	c.LayoutSections["left"]   = "home/left.html"
}

// 分页
func (c *IndexController) Paging () {

	id, _ := strconv.Atoi(c.GetString("id", "2"))
	startPage := (id - 1) * pageNum
	endPage   := id * pageNum

	c.Data["json"] = c.getArtList(startPage, endPage)             // 文章列表
	c.ServeJSON()
}

// 文章标签分类
func (c *IndexController) Category() {
	return
}