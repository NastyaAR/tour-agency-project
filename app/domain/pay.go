package domain

import (
	"github.com/sirupsen/logrus"
)

type PayEvent struct {
	Id    int
	ReqID int
	Sum   int
	State string
}

type IPayAdapter interface {
	SendPaymentRequest(pay *PayEvent, lg *logrus.Logger) error
}
