package contorllercrms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := c.Cookie("userID")
		if err != nil || userId == "" {
			//获取cookie报错或者userId为空则跳转登录界面
			fmt.Println("请先登录")
			c.Redirect(http.StatusTemporaryRedirect, "/crms/login")
		}
		c.Next()
	}

}
