package controller

import (
	"app/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type TourController struct {
	TourService domain.ITourService
	Lg          *logrus.Logger
}

func (tc *TourController) ViewTours(c *gin.Context) {
	tours, _ := tc.TourService.GetTours(0, 10, tc.Lg)

	tmpl, err := template.ParseFiles("./templates/view_tours.tmpl")
	if err != nil {
		fmt.Println("error while parse", err.Error())
	}

	tmpl.Execute(c.Writer, tours)

	c.HTML(http.StatusOK, "view_tours.tmpl", gin.H{
		"title": "View tours",
		"Tours": tours,
	})
}
