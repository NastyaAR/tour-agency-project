package controller

import (
	cli_ui "app/cli-ui"
	"app/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
)

type AccountController struct {
	AccountService domain.IAccountService
	AuthService    domain.IAuthService
}

func (a *AccountController) RegisterClient(lg *logrus.Logger) error {
	acc, err := cli_ui.InputAccount(lg)
	if err != nil {
		lg.Warnf("bad input account")
		return xerrors.Errorf("account controller: register client error: %v", err.Error())
	}

	clnt, err := cli_ui.InputClient(lg)
	if err != nil {
		lg.Warnf("bad input client")
		return xerrors.Errorf("account controller: register client error: %v", err.Error())
	}

	err = a.AccountService.Register(&acc, &clnt, lg)
	if err != nil {
		lg.Warnf("bad register client")
		return xerrors.Errorf("account controller: register client error: %v", err.Error())
	}

	return nil
}

func (a *AccountController) RegisterManager(lg *logrus.Logger) error {
	acc, err := cli_ui.InputAccount(lg)
	if err != nil {
		lg.Warnf("bad input account")
		return xerrors.Errorf("account controller: register manager error: %v", err.Error())
	}

	mngr, err := cli_ui.InputManager(lg)
	if err != nil {
		lg.Warnf("bad input manager")
		return xerrors.Errorf("account controller: register manager error: %v", err.Error())
	}

	err = a.AccountService.Register(&acc, &mngr, lg)
	if err != nil {
		lg.Warnf("bad register manager")
		return xerrors.Errorf("account controller: register manager error: %v", err.Error())
	}

	return nil
}

func (a *AccountController) Login(role string, lg *logrus.Logger) (string, error) {
	acc, err := cli_ui.InputAccountForLogin(lg)
	if err != nil {
		lg.Warnf("bad input account for login")
		return "", xerrors.Errorf("account controller: login error: %v", err.Error())
	}

	if role == domain.AdminUser {
		checkRootAcc, err := a.AccountService.GetByLogin(acc.Login, lg)
		if err != nil {
			lg.Warnf("bad check admin login")
			return "", xerrors.Errorf("account controller: login error: %v", err.Error())
		}
		adminAcc, err := a.AccountService.GetManagerById(checkRootAcc.ID, lg)
		if err != nil {
			lg.Warnf("bad get manager admin acc")
			return "", xerrors.Errorf("account controller: login error: %v", err.Error())
		}
		if adminAcc.Department != "Администрирование" {
			lg.Warnf("bad admin login")
			return "", xerrors.Errorf("account controller: login error: %v", err.Error())
		}
	}

	token, err := a.AccountService.Login(&acc, role, lg)
	if err != nil {
		lg.Warnf("bad login")
		return "", xerrors.Errorf("account controller: login error: %v", err.Error())
	}

	err = a.AuthService.AddToken(acc.Login, token, lg)
	if err != nil {
		lg.Warnf("bad add token")
		return "", xerrors.Errorf("account controller: login error: %v", err.Error())
	}

	return token, nil
}

func (a *AccountController) ProcessAccounts(choice string, lg *logrus.Logger) error {
	var err error

	accessLevels := map[string]string{
		"1": domain.GuestUser,
		"2": domain.AdminUser,
		"3": domain.GuestUser,
		"4": domain.GuestUser,
		"5": domain.GuestUser,
		"0": domain.GuestUser,
	}

	status, err := a.AuthService.CheckAccessRights(domain.TokenString, accessLevels[choice], lg)
	if err != nil && status == domain.ParseTokenError {
		lg.Warnf("tour controller: create error: parse token error")
		return xerrors.Errorf("tour controller: create error: %v", err.Error())
	}
	if err != nil && status == domain.NotAuthorizedError {
		lg.Warnf("tour controller: create error: not authorized")
		return xerrors.Errorf("tour controller: create error: %v", err.Error())
	}

	switch strings.TrimSpace(choice) {
	case "1":
		err = a.RegisterClient(lg)
		if err != nil {
			err = xerrors.New("Ошибка при регистрации клиента")
		}
	case "2":
		err = a.RegisterManager(lg)
		if err != nil {
			err = xerrors.New("Ошибка при регистрации менеджера")
		}
	case "3":
		token, err := a.Login(domain.ClientUser, lg)
		if err != nil {
			err = xerrors.New("Ошибка при входе в личный кабинет клиента")
		} else {
			fmt.Println("\nУспешный вход")
		}
		domain.TokenString = token
	case "4":
		token, err := a.Login(domain.ManagerUser, lg)
		if err != nil {
			err = xerrors.New("Ошибка при входе в личный кабинет менеджера")
		} else {
			fmt.Println("\nУспешный вход")
		}
		domain.TokenString = token
	case "5":
		token, err := a.Login(domain.AdminUser, lg)
		if err != nil {
			err = xerrors.New("Ошибка при входе в личный кабинет администратора")
		} else {
			fmt.Println("\nУспешный вход")
		}
		domain.TokenString = token
	case "6":
		domain.TokenString = ""
	case "0":
		err = nil
	}

	return err
}
