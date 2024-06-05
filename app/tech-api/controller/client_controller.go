package controller

import (
	cli_ui "app/cli-ui"
	comm_parse "app/comm-parse"
	"app/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
)

type ClientController struct {
	ClientService domain.IClientService
	AuthService   domain.IAuthService
}

func (c *ClientController) Create(lg *logrus.Logger) error {
	clntDto, err := cli_ui.InputClient(lg)
	if err != nil {
		lg.Warnf("bad input client")
		return xerrors.Errorf("client controller: create error: %v", err.Error())
	}

	clnt := comm_parse.ToDomainClient(&clntDto)

	_, err = c.ClientService.Create(&clnt, lg)
	if err != nil {
		lg.Warnf("bad create client")
		return xerrors.Errorf("client controller: create error: %v", err.Error())
	}

	return nil
}

func (c *ClientController) Delete(id int, lg *logrus.Logger) error {
	err := c.ClientService.Delete(id, lg)
	if err != nil {
		lg.Warnf("bad delete client")
		return xerrors.Errorf("client controller: delete error: %v", err.Error())
	}
	return nil
}

func (c *ClientController) GetByNameAndSurname(name string, surname string, lg *logrus.Logger) error {
	clnts, err := c.ClientService.GetByNameSurname(name, surname, lg)
	if err != nil {
		lg.Warnf("bad get by name and surname client")
		return xerrors.Errorf("client controller: get by name and surname error: %v", err.Error())
	}
	cli_ui.OutputClients(clnts)
	return nil
}

func (c *ClientController) GetByPhone(phone string, lg *logrus.Logger) error {
	clnts, err := c.ClientService.GetByPhone(phone, lg)
	if err != nil {
		lg.Warnf("bad get by phone client")
		return xerrors.Errorf("client controller: get by phone error: %v", err.Error())
	}
	cli_ui.OutputClients(clnts)
	return nil
}

func (c *ClientController) GetById(id int, lg *logrus.Logger) error {
	clnt, err := c.ClientService.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get by id client")
		return xerrors.Errorf("client controller: get by id error: %v", err.Error())
	}
	cli_ui.PrintClientHeader()
	cli_ui.OutputClient(&clnt)
	return nil
}

func (c *ClientController) GetRequestsStory(lg *logrus.Logger) error {
	id, err := c.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad get id from token")
		return xerrors.Errorf("client controller: get requests story error: %v", err.Error())
	}
	reqs, err := c.ClientService.GetStoryOfRequests(id, lg)
	if err != nil {
		lg.Warnf("bad get story of clients requests")
		return xerrors.Errorf("client controller: get requests story error: %v", err.Error())
	}
	cli_ui.OutputRequests(reqs)
	return nil
}

func (c *ClientController) ProcessClients(choice string, lg *logrus.Logger) error {
	var err error

	accessLevels := map[string]string{
		"1": domain.ManagerUser,
		"3": domain.ManagerUser,
		"2": domain.ManagerUser,
		"4": domain.ClientUser,
		"0": domain.GuestUser,
	}

	status, err := c.AuthService.CheckAccessRights(domain.TokenString, accessLevels[choice], lg)
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
		var name, surname string
		fmt.Print("Введите имя и фамилию клиента: ")
		_, err = fmt.Scanf("%s %s", &name, &surname)
		if err != nil {
			err = xerrors.New("Некорректные имя/фамилия")
			break
		}
		name, surname = strings.TrimSpace(name), strings.TrimSpace(surname)
		err = c.GetByNameAndSurname(name, surname, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении клиентов по имени и фамилии")
		}
	case "2":
		var phone string
		fmt.Printf("Введите телефон клиента: ")
		_, err = fmt.Scanf("%s", &phone)
		if err != nil {
			err = xerrors.New("Некорректный номер телефона")
			break
		}
		phone = strings.TrimSpace(phone)
		err = c.GetByPhone(phone, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении клиентов по номеру телефона")
		}
	case "3":
		var id int
		fmt.Print("Введите id клиента: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = c.GetById(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении клиента по id")
		}
	case "4":
		err = c.GetRequestsStory(lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении истории заявок клиента")
		}
	}
	return nil
}
