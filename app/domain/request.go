package domain

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	Accepted  string = "принята"
	Processed string = "обрабатывается"
	Approved  string = "подтверждена"
	Rejected  string = "отклонена"
	Paid      string = "оплачена"
)

type Request struct {
	ID         int
	TourID     int
	ClntID     int
	MngrID     int
	Status     string
	CreateTime time.Time
	ModifyTime time.Time
	Data       string
}

type IRequestRepo interface {
	GetByStatus(c context.Context, status string, offset int, limit int, lg *logrus.Logger) ([]Request, error)
	Add(c context.Context, req *Request, lg *logrus.Logger) error
	Update(c context.Context, id int, reqCriteria *Request, lg *logrus.Logger) error
	AddTour(c context.Context, id int, tour_id int, lg *logrus.Logger) error
	GetById(c context.Context, id int, lg *logrus.Logger) (Request, error)
	Reject(c context.Context, id int, lg *logrus.Logger) error
	Approve(c context.Context, id int, lg *logrus.Logger) error
	CountFinalCost(c context.Context, id int, lg *logrus.Logger) (int, error)
	GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]Request, error)
	AtomicPay(c context.Context, finalCost int, req *Request, lg *logrus.Logger) error
	UpdateOutbox(c context.Context, id int, lg *logrus.Logger) error
	GetNonSendEvents(c context.Context, limit int, lg *logrus.Logger) ([]PayEvent, error)
	GetRequestsForClient(c context.Context, clnt_id int, offset int, limit int, lg *logrus.Logger) ([]Request, error)
}

type IRequestService interface {
	Create(req *Request, lg *logrus.Logger) error
	GetById(id int, lg *logrus.Logger) (Request, error)
	Update(id int, newState *Request, lg *logrus.Logger) error
	AddTour(id int, tour_id int, mngr_id int, lg *logrus.Logger) error
	GetByStatus(status string, offset int, limit int, lg *logrus.Logger) ([]Request, error)
	Pay(id int, clnt_id int, lg *logrus.Logger) error
	Reject(id int, mngr_id int, lg *logrus.Logger) error
	Approve(id int, mngr_id int, lg *logrus.Logger) error
	GetRequests(offset int, limit int, lg *logrus.Logger) ([]Request, error)
	Paying(done chan bool, frequency time.Duration, lg *logrus.Logger)
	GetRequestsForClient(clnt_id int, offset int, limit int, lg *logrus.Logger) ([]Request, error)
}
