package contorllercrms

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticlesIndex(c *gin.Context){
	c.HTML(http.StatusOK,"articles/index.html",gin.H{
		"title": "文章列表",
	})
}
