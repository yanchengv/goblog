package contorllercrms

import (
	"github.com/gin-gonic/gin"
	pagination2 "goblog/lib/pagination"
	"goblog/models"
	"html/template"
	"net/http"
)

func UserIndex(c *gin.Context) {
	//header的面包屑菜单
	breadcrumb := map[string]string{"m1": "数据管理", "m2": "用户管理", "m2url": "#", "m3": "用户列表"}
	//alertMsg
	//alertMsg := map[string]string{"type": "danger", "msg": "创建成功"}
	pageIndex := 1
	pageSize := 2
	var users []models.User
	var totalCount int64
	models.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&users)
	models.DB.Model(&models.User{}).Count(&totalCount)
	//创建一个分页器，100条数据，每页10条
	pagination := pagination2.Initialize(c.Request, totalCount, pageSize)
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"breadcrumb": breadcrumb,
		//"alertMsg":   alertMsg,
		"paginate":   template.HTML(pagination.Pages()),
		"users":      users,
	})
}
