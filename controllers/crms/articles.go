package contorllercrms

import (
	"github.com/gin-gonic/gin"
	pagination2 "goblog/lib/pagination"
	"goblog/models"
	"html/template"
	"net/http"
	"strconv"
)

func ArticleIndex(c *gin.Context) {
	//header的面包屑菜单
	breadcrumb := map[string]string{"m1": "数据管理", "m2": "文章管理", "m2url": "#", "m3": "文章列表"}
	var articles []models.Article
	var totalCount int64
	pageSize := 20
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	//创建一个分页器，100条数据，每页10条
	models.DB.Model(&models.Article{}).Count(&totalCount)
	pagination := pagination2.Initialize(c.Request, totalCount, pageSize)
	//分页查询记录
	models.DB.Scopes(models.Paginate(page, pageSize)).Find(&articles)
	c.HTML(http.StatusOK, "articles/index.html", gin.H{
		"title":      "文章列表",
		"breadcrumb": breadcrumb,
		"paginate":   template.HTML(pagination.Pages()),
		"articles":   articles,
	})
}

func ArticleNew(c *gin.Context) {
	breadcrumb := map[string]string{"m1": "数据管理", "m2": "文章管理", "m2url": "#", "m3": "创建文章"}
	c.HTML(http.StatusOK, "articles/new.html", gin.H{
		"breadcrumb": breadcrumb,
	})
}

func ArticleCreate(c *gin.Context) {
	title := c.PostForm("title")
	subtitle := c.PostForm("subtitle")
	content := c.PostForm("content")
	article := models.Article{Title: title, Subtitle: subtitle, Content: content}
	models.DB.Create(&article)
	//alertMsg := map[string]string{"type": "success", "msg": "创建成功"}
	// second solution
	//q := url.Values{}
	//q.Set("wage", "123")
	//q.Set("amount", "13123")
	//location := url.URL{Path: "/crms/articles", RawQuery: q.Encode()}
	//c.Redirect(http.StatusFound, location.RequestURI())
	c.Redirect(http.StatusFound, "/crms/articles")
}
