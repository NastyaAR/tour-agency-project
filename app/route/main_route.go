package route

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func MainRouter(router *gin.Engine) {
	//router := gin.Default()
	//router.Static("/css", "./templates/css")
	//router.Static("/static", "./templates/static")
	//router.LoadHTMLGlob("./templates/*.tmpl")
	router.GET("/", controller.Home)
	router.GET("/about", controller.About)
	router.GET("/find", controller.Tours)
}
