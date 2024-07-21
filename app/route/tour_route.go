package route

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func TourRouter(router *gin.Engine, tourCntl *controller.TourController) {
	//router.Static("/css", "./templates/css")
	//router.Static("/static", "./templates/static")
	//router.LoadHTMLGlob("./templates/*.tmpl")
	router.GET("/tours", tourCntl.ViewTours)
}
