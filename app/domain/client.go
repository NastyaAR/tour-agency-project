package domain

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Client struct {
	ID      int
	AccID   int
	Name    string
	Surname string
	Mail    string
	Phone   string
}

type IClientRepo interface {
	Add(c context.Context, clnt *Client, lg *logrus.Logger) (Client, error)
	GetById(c context.Context, id int, lg *logrus.Logger) (Client, error)
	Delete(c context.Context, id int, lg *logrus.Logger) error
	GetByNameSurname(c context.Context, name string, surname string, lg *logrus.Logger) ([]Client, error)
	GetByPhone(c context.Context, phone string, lg *logrus.Logger) ([]Client, error)
	Update(c context.Context, id int, newState *Client, lg *logrus.Logger) error
	GetActiveRequestsByID(c context.Context, id int, lg *logrus.Logger) ([]Request, error)
	GetDoneRequestsByID(c context.Context, id int, lg *logrus.Logger) ([]Request, error)
}

type IClientService interface {
	Create(clnt *Client, lg *logrus.Logger) (Client, error)
	GetById(id int, lg *logrus.Logger) (Client, error)
	Delete(id int, lg *logrus.Logger) error
	GetByNameSurname(name string, surname string, lg *logrus.Logger) ([]Client, error)
	GetByPhone(phone string, lg *logrus.Logger) ([]Client, error)
	Update(id int, newState *Client, lg *logrus.Logger) error
	GetStoryOfRequests(id int, lg *logrus.Logger) ([]Request, error)
}
