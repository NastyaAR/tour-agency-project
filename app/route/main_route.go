package route

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func MainRouter() {
	r := gin.Default()
	r.Static("/css", "./templates/css")
	r.Static("/static", "./templates/static")
	r.LoadHTMLGlob("./templates/*.tmpl")
	r.GET("/", controller.Home)
	r.Run()
}
