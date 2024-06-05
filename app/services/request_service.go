package services

import (
	"app/domain"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type requestService struct {
	RequestRepo     domain.IRequestRepo
	PayAdapter      domain.IPayAdapter
	timeout         time.Duration
	payingDone      chan bool
	payingLimit     int
	payingFrequency time.Duration
}

func CreateNewRequestService(RequestRepo domain.IRequestRepo, PayAdapter domain.IPayAdapter,
	t time.Duration, done chan bool, limit int, pf time.Duration, lg *logrus.Logger) domain.IRequestService {
	newReqService := requestService{RequestRepo: RequestRepo, PayAdapter: PayAdapter, timeout: t,
		payingLimit: limit, payingDone: done, payingFrequency: pf}
	go newReqService.Paying(done, pf, lg)
	return &newReqService
}

func (rs *requestService) Create(req *domain.Request, lg *logrus.Logger) error {
	if req.ClntID < 0 {
		lg.Warnf("bad clnt.id %d", req.ClntID)
		return xerrors.New("request service: create error: bad clnt id < 0")
	}
	if req.MngrID < 0 {
		lg.Warnf("bad mngr.id %d", req.ClntID)
		return xerrors.New("request service: create error: bad mngr id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	err := rs.RequestRepo.Add(ctx, req, lg)
	if err != nil {
		return xerrors.Errorf("request error: create error: %v", err.Error())
	}
	return nil
}

func (rs *requestService) GetById(id int, lg *logrus.Logger) (domain.Request, error) {
	if id < 0 {
		lg.Warnf("bad request id %d", id)
		return domain.Request{}, xerrors.New("request service: getbyid error: bad id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	reqs, err := rs.RequestRepo.GetById(ctx, id, lg)
	if err != nil {
		return domain.Request{}, xerrors.Errorf("request service: getbyid error: %v", err.Error())
	}
	return reqs, nil
}

func (rs *requestService) Update(id int, newState *domain.Request, lg *logrus.Logger) error {
	if newState == nil {
		lg.Warnf("bad newState request = nil")
		return xerrors.New("request service: update error: newState = nil")
	}
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("request service: update error: bad id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	err := rs.RequestRepo.Update(ctx, id, newState, lg)
	if err != nil {
		return xerrors.Errorf("request service: update error: %v", err.Error())
	}
	return nil
}

func (rs *requestService) GetByStatus(status string, offset int, limit int, lg *logrus.Logger) ([]domain.Request, error) {
	if status != domain.Accepted && status != domain.Processed && status != domain.Approved &&
		status != domain.Rejected && status != domain.Paid {
		lg.Warnf("bad request status %s", status)
		return nil, xerrors.New("request service: getbystatus error: bad status")
	}

	if offset < 0 {
		lg.Warnf("bad offset %d", offset)
		return nil, xerrors.New("request repo: getbystatus error: offset < 0")
	}

	if limit < 0 {
		lg.Warnf("bad limit %d", limit)
		return nil, xerrors.New("request repo: getbystatus error: limit < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	reqs, err := rs.RequestRepo.GetByStatus(ctx, status, offset, limit, lg)
	if err != nil {
		return nil, xerrors.Errorf("request service: getbystatus error: %v", err.Error())
	}

	return reqs, nil
}

func (rs *requestService) Pay(id int, clnt_id int, lg *logrus.Logger) error {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("request service: pay error: bad id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	req, err := rs.RequestRepo.GetById(ctx, id, lg)
	if err != nil {
		lg.Warnf("bad get by id request for pay")
		return xerrors.Errorf("request service: pay error: %v", err.Error())
	}

	if req.ClntID != clnt_id {
		lg.Warnf("bad clnt id: request is payed by other client")
		return xerrors.Errorf("request service: pay error: request is payed by other client")
	}

	if req.Status != domain.Approved {
		lg.Warnf("bad request status %s", req.Status)
		return xerrors.New("request service: pay error: bad status")
	}

	finalCost, err := rs.RequestRepo.CountFinalCost(ctx, req.TourID, lg)

	fmt.Printf("\nФинальная стоимость тура: %d (со скидкой)\n", finalCost)

	if err != nil {
		return xerrors.Errorf("request service: pay error: %v", err)
	}

	err = rs.RequestRepo.AtomicPay(ctx, finalCost, &req, lg)
	if err != nil {
		return xerrors.Errorf("request service: pay error: %v", err.Error())
	}

	return nil
}

func (rs *requestService) Paying(done chan bool, frequency time.Duration, lg *logrus.Logger) {
	for {
		select {
		case <-done:
			lg.Warnf("paying gouroutine exited")
			return
		default:
			lg.Info("paying gouroutine served requests")
			ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
			defer cancel()

			reqs, err := rs.RequestRepo.GetNonSendEvents(ctx, rs.payingLimit, lg)
			if err != nil {
				lg.Warnf("pay adapter: paying error: %v", err.Error())
			}

			for _, req := range reqs {
				err = rs.PayAdapter.SendPaymentRequest(&req, lg)
				if err != nil {
					fmt.Printf("error: %v", err.Error())
					lg.Warnf("paying: send error: %v", err.Error())
				} else {
					fmt.Printf("upd start")
					err = rs.RequestRepo.UpdateOutbox(ctx, req.ReqID, lg)
					if err != nil {
						fmt.Printf("error: %v", err.Error())
						lg.Warnf("paying: update outbox error: %v", err.Error())
					}
				}
			}
			time.Sleep(frequency)
		}
	}
}

func (rs *requestService) Reject(id int, mngr_id int, lg *logrus.Logger) error {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("request service: reject error: bad id < 0")
	}

	req, err := rs.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get request")
		return xerrors.Errorf("request service: reject error: %v", err.Error())
	}

	if req.MngrID != mngr_id {
		lg.Warnf("bad mngr id: request is served by other manager")
		return xerrors.Errorf("request service: reject error: request is served by other manager")
	}

	if req.Status == domain.Paid || req.Status == domain.Rejected {
		lg.Warnf("bad status request for rejecting")
		return xerrors.Errorf("request service: reject error: bad status request for rejecting")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	err = rs.RequestRepo.Reject(ctx, id, lg)
	if err != nil {
		return xerrors.Errorf("request error: reject error: %v", err.Error())
	}

	return nil
}

func (rs *requestService) Approve(id int, mngr_id int, lg *logrus.Logger) error {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("request service: approve error: bad id < 0")
	}

	req, err := rs.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get request")
		return xerrors.Errorf("request service: approve error: %v", err.Error())
	}

	if req.MngrID != mngr_id {
		lg.Warnf("bad mngr id: request is served by other manager")
		return xerrors.Errorf("request service: approve error: request is served by other manager")
	}

	if req.Status == domain.Paid || req.Status == domain.Rejected {
		lg.Warnf("bad status request for rejecting")
		return xerrors.Errorf("request service: approve error: bad status request for rejecting")
	}
	if req.TourID == domain.DefaultEmptyValue {
		lg.Warnf("request doesnt link with tour")
		return xerrors.Errorf("request service: approve error: request doesnt link with tour")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	err = rs.RequestRepo.Approve(ctx, id, lg)
	if err != nil {
		return xerrors.Errorf("request error: approve error: %v", err.Error())
	}

	return nil
}

func (rs *requestService) GetRequests(offset int, limit int, lg *logrus.Logger) ([]domain.Request, error) {
	if offset < 0 {
		lg.Warnf("bad offset %d", offset)
		return nil, xerrors.New("request service: getrequests error: bad offset < 0")
	}
	if limit < 0 {
		lg.Warnf("bad limit %d", offset)
		return nil, xerrors.New("request service: getrequests error: bad limit < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	reqs, err := rs.RequestRepo.GetLimit(ctx, offset, limit, lg)
	if err != nil {
		return nil, xerrors.Errorf("requests service: getrequests error: %v", err.Error())
	}

	return reqs, nil
}

func (rs *requestService) AddTour(id int, tour_id int, mngr_id int, lg *logrus.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	req, err := rs.RequestRepo.GetById(ctx, id, lg)
	if err != nil {
		lg.Warnf("bad get by id request")
		return xerrors.Errorf("request service: add tour error: %v", err.Error())
	}

	if req.MngrID != mngr_id {
		lg.Warnf("bad mngr id: request is served by other manager")
		return xerrors.Errorf("request service: addtour error: request is served by other manager")
	}

	if req.TourID != domain.DefaultEmptyValue {
		lg.Warnf("bad add tour: tour exists")
		return xerrors.Errorf("request service: add tour error: tour exists")
	}

	if req.Status == domain.Paid || req.Status == domain.Rejected || req.Status == domain.Approved {
		lg.Warnf("bad status for adding tour")
		return xerrors.Errorf("request service: add tour error: bad status for adding tour")
	}

	err = rs.RequestRepo.AddTour(ctx, id, tour_id, lg)
	if err != nil {
		lg.Warnf("bad adding tour into request")
		return xerrors.Errorf("request serviceL add tour error: %v", err.Error())
	}

	return nil
}

func (rs *requestService) GetRequestsForClient(clnt_id int, offset int, limit int, lg *logrus.Logger) ([]domain.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rs.timeout)
	defer cancel()

	reqs, err := rs.RequestRepo.GetRequestsForClient(ctx, clnt_id, offset, limit, lg)
	if err != nil {
		lg.Warnf("bad get requests: %v", err.Error())
		return nil, xerrors.Errorf("request service: get request for client error: %v", err.Error())
	}
	clntReqs := make([]domain.Request, 0)
	for _, req := range reqs {
		if req.ClntID == clnt_id {
			clntReqs = append(clntReqs, req)
		}
	}

	return clntReqs, nil
}
