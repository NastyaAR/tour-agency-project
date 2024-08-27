package controller

import (
	comm_parse "app/comm-parse"
	"app/domain"
	"app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type TourController struct {
	TourService domain.ITourService
	Pager       pkg.Pager
	Lg          *logrus.Logger
}

var curTours []domain.Tour

//type TourDTO struct {
//	ID             int    `json:"id"`
//	ChillPlace     string `json:"chillPlace"`
//	FromPlace      string `json:"fromPlace"`
//	Date           string `json:"date"`
//	Duration       int    `json:"duration"`
//	Cost           int    `json:"cost"`
//	TouristsNumber int    `json:"touristsNumber"`
//	ChillType      string `json:"chillType"`
//}

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
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	tc.Lg.Info(jsonData)
	var j comm_parse.TourJson

	j, err = comm_parse.FromStringToTourJson(string(jsonData), tc.Lg)
	if err != nil {
		tc.Lg.Warnf("Ошибка десериализации JSON: %v", err)
	}

	criteria, err := j.ToDomainTour(tc.Lg)
	if err != nil {
		tc.Lg.Warnf("to domain tour error: %v", err.Error())
	}

	tours, err := tc.TourService.GetByCriteria(&criteria, 0, pkg.ItemsOnPage, tc.Lg)
	if err != nil {
		tc.Lg.Warnf("get tours by criteria error: %v", err.Error())
	}

	curTours = tours
	c.Redirect(http.StatusSeeOther, "/found?page=1")
}

func (tc *TourController) ViewFoundTours(c *gin.Context) {
	pageNumber, err := tc.Pager.GetPage(c, tc.Lg)
	if err != nil {
		tc.Lg.Errorf("tour controller: findTours error: %v", err.Error())
	}

	c.HTML(http.StatusOK, "view_tours.tmpl", gin.H{
		"title":        "View found tours",
		"Tours":        curTours,
		"current_page": pageNumber,
		"total_pages":  tc.Pager.GetTotalPages(),
	})
}
