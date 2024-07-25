package route

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func TourRouter(router *gin.Engine, tourCntl *controller.TourController) {
	router.GET("/tours", tourCntl.ViewTours)
	router.GET("/find/tours", tourCntl.ViewTours)

}
