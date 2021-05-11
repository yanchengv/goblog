package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"goblog/models"
	admins "goblog/routes"
	"html/template"
	"net/http"
)

//go:embed assets views
var static embed.FS

func main() {
	models.InitDB()
	//user := models.User{Name: "admin1",Email: "admin1@163.com",Password: "123"}
	//models.DB.Create(&user)
	r := gin.Default()
	r.StaticFS("/public", http.FS(static))
	//r.LoadHTMLGlob("views/*/*")
	templ := template.Must(template.New("").ParseFS(static, "views/*/*"))
	r.SetHTMLTemplate(templ)
	//后台路由
	admins.AdminRouters(r)
	r.Run(":9000")
}
