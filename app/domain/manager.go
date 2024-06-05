package domain

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Manager struct {
	ID         int
	AccID      int
	Name       string
	Surname    string
	Department string
}

type Statistics struct {
	ManagerAcc Manager
	Efficiency float32
	SaleSum    int
}

type IManagerRepo interface {
	Add(c context.Context, mngr *Manager, lg *logrus.Logger) (Manager, error)
	Delete(c context.Context, id int, lg *logrus.Logger) error
	GetByNameSurname(c context.Context, name string, surname string, lg *logrus.Logger) ([]Manager, error)
	GetByDepartment(c context.Context, department string, lg *logrus.Logger) ([]Manager, error)
	GetById(c context.Context, id int, lg *logrus.Logger) (Manager, error)
	Update(c context.Context, id int, newState *Manager, lg *logrus.Logger) error
	GetNumberServedRequests(c context.Context, id int, lg *logrus.Logger) (int, error)
	GetAllRequests(c context.Context, id int, lg *logrus.Logger) (int, error)
	GetSumOnPeriod(c context.Context, id int, from time.Time, to time.Time, lg *logrus.Logger) (int, error)
	GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]Manager, error)
}

type IManagerService interface {
	Create(mngr *Manager, lg *logrus.Logger) (Manager, error)
	Delete(id int, lg *logrus.Logger) error
	GetById(id int, lg *logrus.Logger) (Manager, error)
	GetByNameSurname(name string, surname string, lg *logrus.Logger) ([]Manager, error)
	GetByDepartment(department string, lg *logrus.Logger) ([]Manager, error)
	Update(id int, newState *Manager, lg *logrus.Logger) error
	GetStatisticsForManagers(from time.Time, to time.Time, offset int, limit int, lg *logrus.Logger) ([]Statistics, error)
	GetStatisticsForManager(id int, from time.Time, to time.Time, lg *logrus.Logger) (Statistics, error)
}
