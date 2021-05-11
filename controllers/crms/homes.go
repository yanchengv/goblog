package contorllercrms

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeIndex(c *gin.Context){
	c.HTML(http.StatusOK,"homes/index.html",gin.H{
		"title": "主页",
	})
}
