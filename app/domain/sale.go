package domain

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Sale struct {
	ID          int
	Name        string
	ExpiredTime time.Time
	Percent     int
}

type ISaleService interface {
	Create(sale *Sale, lg *logrus.Logger) error
	GetById(id int, lg *logrus.Logger) (Sale, error)
	Update(id int, newState *Sale, lg *logrus.Logger) error
	Delete(id int, lg *logrus.Logger) error
	GetByCriteria(saleCriteria *Sale, offset int, limit int, lg *logrus.Logger) ([]Sale, error)
	GetSales(offset int, limit int, lg *logrus.Logger) ([]Sale, error)
}

type ISaleRepo interface {
	Add(c context.Context, sale *Sale, lg *logrus.Logger) error
	GetById(c context.Context, id int, lg *logrus.Logger) (Sale, error)
	Update(c context.Context, id int, newState *Sale, lg *logrus.Logger) error
	Delete(c context.Context, id int, lg *logrus.Logger) error //возможно, не только по id
	GetByCriteria(c context.Context, offset int, limit int, saleCriteria *Sale, lg *logrus.Logger) ([]Sale, error)
	GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]Sale, error)
}
