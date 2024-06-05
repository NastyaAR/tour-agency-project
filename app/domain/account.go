package domain

import (
	"context"
	"github.com/sirupsen/logrus"
)

const (
	AdminUser   string = "администратор"
	ManagerUser string = "менеджер"
	ClientUser  string = "клиент"
	GuestUser   string = "гость"
)

type Account struct {
	ID       int
	Login    string
	Password string
}

type NewAccountDTO struct {
	ID         int
	Name       string
	Surname    string
	Mail       string
	Phone      string
	Department string
	Role       string
}

type IAccountRepo interface {
	Add(c context.Context, acc *Account, newAccDTO *NewAccountDTO, lg *logrus.Logger) error
	Delete(c context.Context, id int, lg *logrus.Logger) error
	GetByLogin(c context.Context, login string, lg *logrus.Logger) (Account, error)
	AddClient(c context.Context, dto *NewAccountDTO, lg *logrus.Logger) (int, error)
	AddManager(c context.Context, dto *NewAccountDTO, lg *logrus.Logger) (int, error)
	GetClientById(c context.Context, accId int, lg *logrus.Logger) (Client, error)
	GetManagerById(c context.Context, accId int, lg *logrus.Logger) (Manager, error)
}

type IAccountService interface {
	Register(acc *Account, newAccDTO *NewAccountDTO, lg *logrus.Logger) error
	Login(acc *Account, role string, lg *logrus.Logger) (string, error)
	GetByLogin(login string, lg *logrus.Logger) (Account, error)
	GetClientById(id int, lg *logrus.Logger) (Client, error)
	GetManagerById(id int, lg *logrus.Logger) (Manager, error)
}
