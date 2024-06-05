package services

import (
	"app/domain"
	"app/pkg"
	"context"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
	"time"
)

type accountService struct {
	accountRepo domain.IAccountRepo
	timeout     time.Duration
}

func CreateNewAccountService(accountRepo domain.IAccountRepo, t time.Duration) domain.IAccountService {
	return &accountService{accountRepo: accountRepo,
		timeout: t,
	}
}

func (as *accountService) Register(acc *domain.Account, newAccDTO *domain.NewAccountDTO, lg *logrus.Logger) error {
	if acc == nil {
		lg.Warnf("bad acc = nil")
		return xerrors.New("account service: register error: bad acc nil")
	}

	if newAccDTO == nil {
		lg.Warnf("bad newAccDTO = nil")
		return xerrors.New("account service: register error: bad newAccDTO nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), as.timeout)
	defer cancel()

	_, err := as.accountRepo.GetByLogin(ctx, acc.Login, lg)
	if err != nil {
		return xerrors.Errorf("account service: register error: %v", err.Error())
	}

	pswd := []byte(acc.Password)
	crPswd, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
	if err != nil {
		lg.Warnf("bad crypting password")
		return xerrors.New("account service: reqister error: bcrypt error")
	}
	acc.Password = string(crPswd)
	err = as.accountRepo.Add(ctx, acc, newAccDTO, lg)
	if err != nil {
		return xerrors.Errorf("account service: register error: %v", err.Error())
	}

	lg.Info("account service: successfully register")

	return nil
}

func (as *accountService) Login(acc *domain.Account, role string, lg *logrus.Logger) (string, error) {
	if acc == nil {
		lg.Warnf("bad acc = nil")
		return "", xerrors.New("account service: login error: bad acc = nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), as.timeout)
	defer cancel()

	user, err := as.accountRepo.GetByLogin(ctx, acc.Login, lg)
	if err != nil {
		return "", xerrors.Errorf("account service: login error: %v", err.Error())
	}
	receivedPswd := []byte(acc.Password)
	passwordDB := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordDB, receivedPswd)
	if err != nil {
		lg.Warnf("bcrypt error")
		return "", xerrors.Errorf("account service: login error: %v", err.Error())
	}

	var clnt domain.Client
	var mngr domain.Manager
	var userId int
	switch {
	case role == domain.ClientUser:
		clnt, err = as.GetClientById(user.ID, lg)
		if err != nil {
			lg.Warnf("get client by account id error")
			return "", xerrors.Errorf("account service: login error: %v", err.Error())
		}
		userId = clnt.ID
	case role == domain.ManagerUser || role == domain.AdminUser:
		mngr, err = as.GetManagerById(user.ID, lg)
		if err != nil {
			lg.Warnf("get client by account id error")
			return "", xerrors.Errorf("account service: login error: %v", err.Error())
		}
		userId = mngr.ID
	}

	token, err := pkg.GenerateJWTToken(userId, role)
	if err != nil {
		lg.Warnf("generate jwt error")
		return "", xerrors.Errorf("account service: login error: %v", err.Error())
	}

	return token, nil
}

func (as *accountService) GetByLogin(login string, lg *logrus.Logger) (domain.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), as.timeout)
	defer cancel()

	acc, err := as.accountRepo.GetByLogin(ctx, login, lg)
	if err != nil {
		return domain.Account{}, xerrors.Errorf("account service: getbylogin error: %s", err.Error())
	}
	return acc, err
}

func (as *accountService) GetClientById(accId int, lg *logrus.Logger) (domain.Client, error) {
	if accId < 0 {
		lg.Warnf("bad accId %d", accId)
		return domain.Client{},
			xerrors.New("account service: getclientbyid error: bad accId < 0")
	}
	ctx, cancel := context.WithTimeout(context.Background(), as.timeout)
	defer cancel()

	clnt, err := as.accountRepo.GetClientById(ctx, accId, lg)
	if err != nil {
		return domain.Client{}, xerrors.Errorf("account service: getclientbyuserid error: %v", err.Error())
	}
	return clnt, nil
}

func (as *accountService) GetManagerById(userID int, lg *logrus.Logger) (domain.Manager, error) {
	if userID < 0 {
		lg.Warnf("bad userID %d", userID)
		return domain.Manager{},
			xerrors.New("account service: getmanagerbyuserid error: bad userID < 0")
	}
	ctx, cancel := context.WithTimeout(context.Background(), as.timeout)
	defer cancel()

	mngr, err := as.accountRepo.GetManagerById(ctx, userID, lg)
	if err != nil {
		return domain.Manager{}, xerrors.Errorf("account service: getmanagerbyuserid error: %v", err.Error())
	}
	return mngr, nil
}
