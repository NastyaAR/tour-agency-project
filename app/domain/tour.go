package domain

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	Elite   string = "элитный"
	City    string = "городской"
	Beach   string = "пляжный"
	Active  string = "активный"
	Unusual string = "необычный"
)

type HotTourDto struct {
	ID             int
	ChillPlace     string
	FromPlace      string
	Date           time.Time
	Duration       int
	Cost           int
	TouristsNumber int
	ChillType      string
	SaleName       string
	SalePercemt    int
}

type Tour struct {
	ID             int
	ChillPlace     string
	FromPlace      string
	Date           time.Time
	Duration       int
	Cost           int
	TouristsNumber int
	ChillType      string
}

const (
	DefaultEmptyValue = -1
)

type ITourRepo interface {
	GetByCriteria(c context.Context, offset int, limit int, criteria *Tour, lg *logrus.Logger) ([]Tour, error)
	GetById(c context.Context, id int, lg *logrus.Logger) (Tour, error)
	Add(c context.Context, tour *Tour, lg *logrus.Logger) error
	Delete(c context.Context, id int, lg *logrus.Logger) error //возможно не по id только
	Update(c context.Context, id int, newState *Tour, lg *logrus.Logger) error
	UpdateSale(c context.Context, id int, newSale *Sale, lg *logrus.Logger) error
	GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]Tour, error)
	GetHotTours(c context.Context, offset int, limit int, lg *logrus.Logger) ([]HotTourDto, error)
	GetNumberOfTours(c context.Context, lg *logrus.Logger) (int, error)
}

type ITourService interface {
	Create(tour *Tour, lg *logrus.Logger) error
	GetById(id int, lg *logrus.Logger) (Tour, error)
	Update(id int, newState *Tour, lg *logrus.Logger) error
	SetSale(id int, newSale *Sale, lg *logrus.Logger) error
	Delete(id int, lg *logrus.Logger) error //возможно не по id только
	GetByCriteria(criteria *Tour, offset int, limit int, lg *logrus.Logger) ([]Tour, error)
	GetTours(offset int, limit int, lg *logrus.Logger) ([]Tour, error)
	GetHotTours(offset int, limit int, lg *logrus.Logger) ([]HotTourDto, error)
	GetNumberOfTours(lg *logrus.Logger) (int, error)
}
