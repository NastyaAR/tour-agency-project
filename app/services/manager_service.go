package services

import (
	"app/domain"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type managerService struct {
	managerRepo domain.IManagerRepo
	timeout     time.Duration
}

func CreateNewManagerService(managerRepo domain.IManagerRepo, t time.Duration) domain.IManagerService {
	return &managerService{managerRepo: managerRepo, timeout: t}
}

func (ms *managerService) Create(mngr *domain.Manager, lg *logrus.Logger) (domain.Manager, error) {
	if mngr.ID < 0 {
		lg.Warnf("bad mngr.ID %d", mngr.ID)
		return domain.Manager{}, xerrors.New("Некорректный id менеджера!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	addedMngr, err := ms.managerRepo.Add(ctx, mngr, lg)
	if err != nil {
		return domain.Manager{}, xerrors.Errorf("manager service: create error: %v", err.Error())
	}
	return addedMngr, nil
}

func (ms *managerService) Delete(id int, lg *logrus.Logger) error {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("manager service: delete error: bad id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	err := ms.managerRepo.Delete(ctx, id, lg)
	if err != nil {
		return xerrors.Errorf("manager service: delete error: %v", err.Error())
	}
	return nil
}

func (ms *managerService) GetByNameSurname(name string, surname string, lg *logrus.Logger) ([]domain.Manager, error) {
	if name == "" && surname == "" {
		lg.Warnf("bad name=%s or surname=%s", name, surname)
		return nil, xerrors.New("manager service: getbynamesurname error: empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	mngrs, err := ms.managerRepo.GetByNameSurname(ctx, name, surname, lg)
	if err != nil {
		return nil, xerrors.Errorf("manager service: getbynamesurname error: %v", err.Error())
	}
	return mngrs, nil
}

func (ms *managerService) GetByDepartment(department string, lg *logrus.Logger) ([]domain.Manager, error) {
	if department == "" {
		lg.Warnf("bad department = nil")
		return nil, xerrors.New("manager service: getbydepartment error: dep = nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	mngrs, err := ms.managerRepo.GetByDepartment(ctx, department, lg)
	if err != nil {
		return nil, xerrors.Errorf("manager service: getbydepartment error: %v", err.Error())
	}
	return mngrs, nil
}

func (ms *managerService) Update(id int, newState *domain.Manager, lg *logrus.Logger) error {
	if newState == nil {
		lg.Warnf("bad newState manager = nil")
		return xerrors.New("manager service: update error: newState = nil")
	}

	if id < 0 {
		lg.Warnf("bad id %d", id)
		return xerrors.New("manager service: update error: bad id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	err := ms.managerRepo.Update(ctx, id, newState, lg)
	if err != nil {
		return xerrors.Errorf("manager service: update error: %v", err.Error())
	}
	return nil
}

func (ms *managerService) GetStatisticsForManagers(from time.Time, to time.Time, offset int, limit int, lg *logrus.Logger) ([]domain.Statistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	mngrs, err := ms.managerRepo.GetLimit(ctx, offset, limit, lg)
	if err != nil {
		return nil, xerrors.Errorf("manager service: getstatistics error: %v", err.Error())
	}

	statistics := make([]domain.Statistics, 0)
	for _, mngr := range mngrs {
		stat, err := ms.GetStatisticsForManager(mngr.ID, from, to, lg)
		fmt.Println(stat)
		if err == nil {
			statistics = append(statistics, stat)
		}
	}

	return statistics, nil
}

func (ms *managerService) GetStatisticsForManager(id int, from time.Time, to time.Time, lg *logrus.Logger) (domain.Statistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	mngr, err := ms.managerRepo.GetById(ctx, id, lg)
	if err != nil {
		return domain.Statistics{},
			xerrors.Errorf("manager service: getonestatistics error: %v", err.Error())
	}

	reqs, err := ms.managerRepo.GetAllRequests(ctx, id, lg)
	if err != nil {
		return domain.Statistics{},
			xerrors.Errorf("manager service: getonestatistics error: %v", err.Error())
	}

	servReqs, err := ms.managerRepo.GetNumberServedRequests(ctx, id, lg)
	if err != nil {
		return domain.Statistics{},
			xerrors.Errorf("manager service: getonestatistics error: %v", err.Error())
	}

	if servReqs == 0 {
		lg.Warnf("bad number of served requests %d", servReqs)
		return domain.Statistics{},
			xerrors.New("manager service: getonestatistics error: bad number")
	}

	sum, err := ms.managerRepo.GetSumOnPeriod(ctx, id, from, to, lg)
	if err != nil {
		return domain.Statistics{},
			xerrors.Errorf("manager service: getonestatistics error: %v", err.Error())
	}

	var eff float32
	eff = float32(servReqs) / float32(reqs) * 100
	stat := domain.Statistics{
		ManagerAcc: mngr,
		Efficiency: eff,
		SaleSum:    sum,
	}
	return stat, nil
}

func (ms *managerService) GetById(id int, lg *logrus.Logger) (domain.Manager, error) {
	if id < 0 {
		lg.Warnf("bad id %d", id)
		return domain.Manager{}, xerrors.New("manager service: getbyid error: bad id < 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.timeout)
	defer cancel()

	mngr, err := ms.managerRepo.GetById(ctx, id, lg)
	if err != nil {
		lg.Warnf("manager repo get by id error")
		return domain.Manager{}, xerrors.Errorf("manager service: getbyid error: %v", err.Error())
	}

	return mngr, nil
}
