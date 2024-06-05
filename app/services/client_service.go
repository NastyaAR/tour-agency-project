package services

import (
	"app/domain"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type clientService struct {
	clientRepo domain.IClientRepo
	timeout    time.Duration
}

func CreateNewClientService(clientRepo domain.IClientRepo, t time.Duration) domain.IClientService {
	return &clientService{clientRepo: clientRepo, timeout: t}
}

func (cs *clientService) Create(clnt *domain.Client, lg *logrus.Logger) (domain.Client, error) {
	if clnt == nil {
		lg.Warnf("bad clnt = nil")
		return domain.Client{}, errors.New("client service: create error: bad clnt = nil")
	}
	if clnt.ID < 0 {
		lg.Warnf("bad clnt.id %d", clnt.ID)
		return domain.Client{}, xerrors.New("client service: create error: bad id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), cs.timeout)
	defer cancel()

	addedClnt, err := cs.clientRepo.Add(ctx, clnt, lg)
	if err != nil {
		return domain.Client{}, xerrors.Errorf("client service: create error: %v", err.Error())
	}
	return addedClnt, nil
}

func (cs *clientService) Delete(id int, lg *logrus.Logger) error {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("client service: delete error: bad id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), cs.timeout)
	defer cancel()

	err := cs.clientRepo.Delete(ctx, id, lg)
	if err != nil {
		return xerrors.Errorf("client service: delete error: %v", err.Error())
	}
	return nil
}

func (cs *clientService) GetByNameSurname(name string, surname string, lg *logrus.Logger) ([]domain.Client, error) {
	if name == "" || surname == "" {
		lg.Warnf("bad name=%s or surname=%s", name, surname)
		return nil, xerrors.New("client service error: getbynamesurname error: bad name or surname")
	}

	ctx, cancel := context.WithTimeout(context.Background(), cs.timeout)
	defer cancel()

	clnts, err := cs.clientRepo.GetByNameSurname(ctx, name, surname, lg)
	if err != nil {
		return nil, xerrors.Errorf("client service: getbynamesurname error: %v", err.Error())
	}
	return clnts, nil
}

func (cs *clientService) GetByPhone(phone string, lg *logrus.Logger) ([]domain.Client, error) {
	if phone == "" {
		lg.Warnf("bad phone", phone)
		return nil, xerrors.New("client service error: getbyphone error: bad phone")
	}

	ctx, cancel := context.WithTimeout(context.Background(), cs.timeout)
	defer cancel()

	clnts, err := cs.clientRepo.GetByPhone(ctx, phone, lg)
	if err != nil {
		return nil, xerrors.Errorf("client service: getbyphone error: %v", err.Error())
	}

	return clnts, nil
}

func (cs *clientService) Update(id int, newState *domain.Client, lg *logrus.Logger) error {
	if newState == nil {
		lg.Warnf("bad newState for updating client = nil")
		return xerrors.New("client service: update error: bad newState = nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), cs.timeout)
	defer cancel()

	err := cs.clientRepo.Update(ctx, id, newState, lg)
	if err != nil {
		return xerrors.Errorf("client service: update error: %v", err.Error())
	}
	return nil
}

func (cs *clientService) GetStoryOfRequests(id int, lg *logrus.Logger) ([]domain.Request, error) {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return nil, xerrors.New("client service: getstoryofrequests error: bad id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), cs.timeout)
	defer cancel()

	reqs, err := cs.clientRepo.GetActiveRequestsByID(ctx, id, lg)
	if err != nil {
		return nil, xerrors.Errorf("client service: getstoryofrequests error: %v", err.Error())
	}

	doneReqs, err := cs.clientRepo.GetDoneRequestsByID(ctx, id, lg)
	if err != nil {
		return reqs, xerrors.Errorf("client service: getstoryofrequests error: %v", err.Error())
	}
	allReqs := append(doneReqs, reqs...)
	return allReqs, nil
}

func (cs *clientService) GetById(id int, lg *logrus.Logger) (domain.Client, error) {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return domain.Client{}, xerrors.New("client service: getbyid error: id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), cs.timeout)
	defer cancel()

	clnt, err := cs.clientRepo.GetById(ctx, id, lg)
	if err != nil {
		return domain.Client{}, xerrors.Errorf("client service: getbyid error: %v", err.Error())
	}
	return clnt, nil
}
