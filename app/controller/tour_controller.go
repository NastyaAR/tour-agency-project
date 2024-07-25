package controller

import (
	"app/domain"
	"app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TourController struct {
	TourService domain.ITourService
	Pager       pkg.Pager
	Lg          *logrus.Logger
}

func (tc *TourController) ViewTours(c *gin.Context) {
	pageNumber, err := tc.Pager.GetPage(c, tc.Lg)
	if err != nil {
		tc.Lg.Errorf("tour controller: viewTours error: %v", err.Error())
	}
	tours, _ := tc.TourService.GetTours((pageNumber-1)*pkg.ItemsOnPage, pkg.ItemsOnPage, tc.Lg)

	c.HTML(http.StatusOK, "view_tours.tmpl", gin.H{
		"title":        "View tours",
		"Tours":        tours,
		"current_page": pageNumber,
		"total_pages":  tc.Pager.GetTotalPages(),
	})
}

func (tc *TourController) FindTours(c *gin.Context) {

}
