package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"math"
	"strconv"
)

const DefaultPageNumber = 1
const ItemsOnPage = 10

type OutputByPages interface {
	GetNumber(lg *logrus.Logger) (int, error)
}

type Pager interface {
	GetPage(c *gin.Context, lg *logrus.Logger) (int, error)
	GetNumberOfPages(lg *logrus.Logger) (int, error)
	GetCurPage() int
	GetTotalPages() int
}

type GinPager struct {
	curPage    int
	totalPages int
	outputObj  OutputByPages
}

func (gp *GinPager) GetPage(c *gin.Context, lg *logrus.Logger) (int, error) {
	pageStrNumber := c.Query("page")
	if pageStrNumber == "" {
		lg.Errorf("pager: getPage error: empty param string")
		return gp.curPage, xerrors.Errorf("pager: getPage error: empty param string")
	}
	pageNumber, err := strconv.Atoi(pageStrNumber)
	if err != nil {
		lg.Errorf("pager: getPage error: atoi error: %v", err.Error())
		return gp.curPage, xerrors.Errorf("pager: getPage error: atoi error: %v", err.Error())
	}

	if pageNumber < 1 {
		lg.Errorf("pager: getPage error: negative page number")
		return gp.curPage, xerrors.Errorf("pager: getPage error: negative page number")
	}

	if pageNumber > gp.totalPages {
		lg.Errorf("pager: getPage error: out of range page number")
		return gp.curPage, xerrors.Errorf("pager: getPage error: out of range page number")
	}

	return pageNumber, nil
}

func (gp *GinPager) GetNumberOfPages(lg *logrus.Logger) (int, error) {
	totalNumber, err := gp.outputObj.GetNumber(lg)
	lg.Info(totalNumber)
	totalPages := int(math.Ceil(float64(totalNumber) / ItemsOnPage))
	lg.Info(totalPages)
	if err != nil {
		lg.Errorf("pager: getNumberOfPages error: %v", err.Error())
		return DefaultPageNumber, xerrors.Errorf("pager: getNumberOfPages error: %v", err.Error())
	}

	return totalPages, nil
}

func CreateGinPager(outputByPages OutputByPages, lg *logrus.Logger) (Pager, error) {
	pager := GinPager{
		curPage:    1,
		totalPages: DefaultPageNumber,
		outputObj:  outputByPages,
	}
	var err error
	pager.totalPages, err = pager.GetNumberOfPages(lg)
	if err != nil {
		lg.Errorf("pager: createGinPager error: %v", err.Error())
		return nil, xerrors.Errorf("pager: createGinPager error: %v", err.Error())
	}

	return &pager, nil
}

func (gp *GinPager) GetCurPage() int {
	return gp.curPage
}

func (gp *GinPager) GetTotalPages() int {
	return gp.totalPages
}

func (gp *GinPager) SetTotalPages(tp int) {
	gp.totalPages = tp
}
