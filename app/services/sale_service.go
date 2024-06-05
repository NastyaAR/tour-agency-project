package services

import (
	"app/domain"
	"context"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type saleService struct {
	saleRepo domain.ISaleRepo
	timeout  time.Duration
}

func NewSaleService(saleRepo domain.ISaleRepo, t time.Duration) domain.ISaleService {
	return &saleService{saleRepo: saleRepo, timeout: t}
}

func (sr *saleService) Create(sale *domain.Sale, lg *logrus.Logger) error {
	if sale == nil {
		lg.Warnf("bad sale = nil")
		return xerrors.New("sale service: create error: sale = nil")
	}

	if sale.Percent < 0 || sale.Percent > 100 {
		lg.Warnf("bad sale.Percent %d", sale.Percent)
		return xerrors.New("sale service: create error: sale.Percent < 0 || > 100")
	}

	ctx, cancel := context.WithTimeout(context.Background(), sr.timeout)
	defer cancel()

	err := sr.saleRepo.Add(ctx, sale, lg)
	if err != nil {
		return xerrors.Errorf("sale service: create error: %v", err.Error())
	}
	return nil
}

func (sr *saleService) GetById(id int, lg *logrus.Logger) (domain.Sale, error) {
	if id < 0 {
		lg.Warnf("bad sale id %d", id)
		return domain.Sale{}, xerrors.New("sale service: getbyid error: bad id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), sr.timeout)
	defer cancel()

	sales, err := sr.saleRepo.GetById(ctx, id, lg)
	if err != nil {
		return domain.Sale{}, xerrors.Errorf("sale service: getbyid error: %v", err.Error())
	}
	return sales, nil
}

func (sr *saleService) Update(id int, newState *domain.Sale, lg *logrus.Logger) error {
	if newState == nil {
		lg.Warnf("bad newState of sale = nil")
		return xerrors.New("sale service: update error: bad newState = nil")
	}
	if newState.Percent < 0 || newState.Percent > 100 {
		lg.Warnf("bad newState.Percent %d", newState.Percent)
		return xerrors.New("sale service: update error: bad newState.Percent")
	}
	if id < 0 {
		lg.Warnf("bad id for updating sale %d", id)
		return xerrors.New("sale service: update error: bad id")
	}
	if newState.ExpiredTime.Before(time.Now()) {
		lg.Warnf("bad expired time")
		return xerrors.New("sale service: update error: bad expired time")
	}

	ctx, cancel := context.WithTimeout(context.Background(), sr.timeout)
	defer cancel()

	err := sr.saleRepo.Update(ctx, id, newState, lg)
	if err != nil {
		return xerrors.Errorf("sale service: update error: %v", err.Error())
	}
	return nil
}

func (sr *saleService) Delete(id int, lg *logrus.Logger) error {
	if id < 0 {
		lg.Warnf("bad id for deleting sale %d", id)
		return xerrors.New("sale service: delete error: bad id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), sr.timeout)
	defer cancel()

	err := sr.saleRepo.Delete(ctx, id, lg)
	if err != nil {
		return xerrors.Errorf("sale service: delete error: %v", err.Error())
	}
	return nil
}

func (sr *saleService) GetByCriteria(saleCriteria *domain.Sale, offset int, limit int, lg *logrus.Logger) ([]domain.Sale, error) {
	if saleCriteria == nil {
		lg.Warnf("bad saleCriteria = nil")
		return nil, xerrors.New("sale service: getbycriteria error: bad saleCriteria = nil")
	}

	if offset < 0 {
		lg.Warnf("bad offset %d", offset)
		return nil, xerrors.New("sale service: getbycriteria error: bad offset")
	}

	if limit < 0 {
		lg.Warnf("bad limit %d", limit)
		return nil, xerrors.New("sale service: getbycriteria error: bad limit")
	}

	ctx, cancel := context.WithTimeout(context.Background(), sr.timeout)
	defer cancel()

	sales, err := sr.saleRepo.GetByCriteria(ctx, offset, limit, saleCriteria, lg)

	if err != nil {
		return nil, xerrors.Errorf("sale service: getbycriteria error: %v", err.Error())
	}
	return sales, nil
}

func (sr *saleService) GetSales(offset int, limit int, lg *logrus.Logger) ([]domain.Sale, error) {
	if offset < 0 {
		lg.Warnf("bad offset %d", offset)
		return nil, xerrors.New("sale service: getsales error: bad offset < 0")
	}
	if limit < 0 {
		lg.Warnf("bad limit %d", offset)
		return nil, xerrors.New("sale service: getsales error: bad limit < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), sr.timeout)
	defer cancel()

	sales, err := sr.saleRepo.GetLimit(ctx, offset, limit, lg)
	if err != nil {
		return nil, xerrors.Errorf("sale service: getsales error: %v", err.Error())
	}

	return sales, nil
}
