package contorllercrms

import (
	"github.com/gin-gonic/gin"
	pagination "goblog/lib"
	"goblog/models"
	"html/template"
	"net/http"
)

func UserIndex(c *gin.Context) {
	pageIndex := 1
	pageSize := 2
	var users []models.User
	var totalCount int64
	models.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&users)
	models.DB.Model(&models.User{}).Count(&totalCount)
	//创建一个分页器，100条数据，每页10条
	pagination := pagination.Initialize(c.Request, totalCount, pageSize)
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"paginate": template.HTML(pagination.Pages()),
		"users": users,
	})
}
