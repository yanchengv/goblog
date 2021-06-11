package admins

import (
	"github.com/gin-gonic/gin"
	contorllercrms "goblog/controllers/crms"
)

//后台路由
func AdminRouters(r *gin.Engine) {
	r.GET("/", contorllercrms.Login)
	crms := r.Group("/crms")

	crms.GET("/login", contorllercrms.LoginIndex)
	crms.Use(contorllercrms.AuthUser())
	{
		crms.POST("/login", contorllercrms.Login)

		crms.GET("/homes", contorllercrms.HomeIndex)
		users := crms.Group("/users")
		{
			users.GET("", contorllercrms.UserIndex)
			users.POST("/destroy", contorllercrms.UserDestroy)
		}
		articles := crms.Group("/articles")
		{
			articles.GET("", contorllercrms.ArticleIndex)
			articles.GET("/new", contorllercrms.ArticleNew)
			articles.POST("/create", contorllercrms.ArticleCreate)
			articles.POST("/destroy", contorllercrms.ArticleDestroy)
		}

	}

}
