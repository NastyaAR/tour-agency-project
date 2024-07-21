package services

import (
	"app/domain"
	"context"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type tourService struct {
	tourRepository domain.ITourRepo
	timeout        time.Duration
}

func NewTourService(TourRepo domain.ITourRepo, t time.Duration) domain.ITourService {
	return &tourService{
		tourRepository: TourRepo,
		timeout:        t,
	}
}

func (ts *tourService) Create(tour *domain.Tour, lg *logrus.Logger) error {
	if tour.Cost < 0 {
		lg.Warnf("bad tour.Cost %d", tour.Cost)
		return xerrors.New("tour service: create error: bad cost < 0")
	}

	if tour.TouristsNumber < 1 {
		lg.Warnf("bad tour.TouristsNumber %d", tour.TouristsNumber)
		return xerrors.New("tour service: create error: bad touristsNumber < 1")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	err := ts.tourRepository.Add(ctx, tour, lg)
	if err != nil {
		return xerrors.Errorf("tour service: create error: %v", err.Error())
	}
	return nil
}

func (ts *tourService) GetById(id int, lg *logrus.Logger) (domain.Tour, error) {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return domain.Tour{}, xerrors.New("tour service: getbyid error: id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	tours, err := ts.tourRepository.GetById(ctx, id, lg)
	if err != nil {
		return domain.Tour{}, xerrors.Errorf("tour service: getbyid error: %v", err.Error())
	}
	return tours, nil
}

func (ts *tourService) Update(id int, newState *domain.Tour, lg *logrus.Logger) error {
	if newState == nil {
		lg.Warnf("bad newState of tour - nil")
		return xerrors.New("tour service: update error: newState = nil")
	}
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("tour service: update error: bad id < 0")
	}
	if newState.Date.Before(time.Now()) {
		lg.Warnf("bad tour date")
		return xerrors.New("tour service: update error: bad tour date")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	err := ts.tourRepository.Update(ctx, id, newState, lg)
	if err != nil {
		return xerrors.Errorf("tour service: update error: %v", err.Error())
	}
	return nil
}

func (ts *tourService) SetSale(id int, newSale *domain.Sale, lg *logrus.Logger) error {
	if newSale == nil {
		lg.Warnf("bad newSale - nil")
		return xerrors.New("tour service: setsale error: newSale = nil!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	err := ts.tourRepository.UpdateSale(ctx, id, newSale, lg)
	if err != nil {
		return xerrors.Errorf("tour service: setsale error: %v", err.Error())
	}
	return nil
}

func (ts *tourService) Delete(id int, lg *logrus.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	err := ts.tourRepository.Delete(ctx, id, lg)
	if err != nil {
		return xerrors.Errorf("tour service: delete error: %v", err.Error())
	}
	return nil
}

func (ts *tourService) GetByCriteria(criteria *domain.Tour, offset int, limit int, lg *logrus.Logger) ([]domain.Tour, error) {
	if criteria == nil {
		lg.Warnf("bad tour criteria = nil")
		return nil, xerrors.New("tour service: getbycritaria error: bad criteria = nil")
	}

	if offset < 0 {
		lg.Warnf("bad offset %d", offset)
		return nil, xerrors.New("tour service: getbycriteria error: bad offset")
	}

	if limit < 0 {
		lg.Warnf("bad limit %d", limit)
		return nil, xerrors.New("tour service: getbycriteria error: bad limit")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	tours, err := ts.tourRepository.GetByCriteria(ctx, offset, limit, criteria, lg)
	if err != nil {
		return nil, xerrors.Errorf("tour service: getbycriteria error: %v", err.Error())
	}
	return tours, nil
}

func (ts *tourService) GetTours(offset int, limit int, lg *logrus.Logger) ([]domain.Tour, error) {
	if offset < 0 {
		lg.Warnf("bad offset %d", offset)
		return nil, xerrors.New("tour service: gettours error: bad offset < 0")
	}
	if limit < 0 {
		lg.Warnf("bad limit %d", offset)
		return nil, xerrors.New("tour service: gettours error: bad limit < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	tours, err := ts.tourRepository.GetLimit(ctx, offset, limit, lg)
	if err != nil {
		return nil, xerrors.Errorf("tour service: gettours error: %v", err.Error())
	}

	return tours, nil
}

func (ts *tourService) GetHotTours(offset int, limit int, lg *logrus.Logger) ([]domain.HotTourDto, error) {
	if offset < 0 {
		lg.Warnf("bad offset %d", offset)
		return nil, xerrors.New("tour service: gethottours error: bad offset < 0")
	}
	if limit < 0 {
		lg.Warnf("bad limit %d", offset)
		return nil, xerrors.New("tour service: gethottours error: bad limit < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	tours, err := ts.tourRepository.GetHotTours(ctx, offset, limit, lg)
	if err != nil {
		return nil, xerrors.Errorf("tour service: gethottours error: %v", err.Error())
	}

	return tours, nil
}

func (ts *tourService) GetNumberOfTours(lg *logrus.Logger) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ts.timeout)
	defer cancel()

	number, err := ts.tourRepository.GetNumberOfTours(ctx, lg)
	if err != nil {
		lg.Errorf("bad getnumberoftours")
		return domain.DefaultEmptyValue, xerrors.Errorf("tour service: getnumberoftours error: %v", err.Error())
	}

	return number, nil
}
