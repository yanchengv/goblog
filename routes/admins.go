package admins

import (
	"github.com/gin-gonic/gin"
	contorllercrms "goblog/controllers/crms"
)

//后台路由
func AdminRouters(r *gin.Engine) {
	crms := r.Group("/crms")
	crms.GET("/login", contorllercrms.LoginIndex)
	crms.Use(contorllercrms.AuthUser())
	{
		crms.POST("/login",contorllercrms.Login)
		crms.GET("/users", contorllercrms.UserIndex)
		crms.GET("/homes", contorllercrms.HomeIndex)
		articles := crms.Group("/articles")
		{
			articles.GET("", contorllercrms.ArticlesIndex)
		}
	}

}
