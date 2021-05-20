package contorllercrms

import (
	"github.com/gin-gonic/gin"
	pagination "goblog/lib"
	"goblog/models"
	"html/template"
	"net/http"
)

func ArticleIndex(c *gin.Context) {
	//header的面包屑菜单
	breadcrumb := map[string]string{"m1": "数据管理", "m2": "文章管理", "m2url": "#", "m3": "文章列表"}
	pageIndex := 1
	pageSize := 2
	var totalCount int64
	var articles []models.Article
	models.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&articles)
	pagination := pagination.Initialize(c.Request, totalCount, pageSize)
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
