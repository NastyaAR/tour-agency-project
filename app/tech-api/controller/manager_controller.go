package controller

import (
	cli_ui "app/cli-ui"
	comm_parse "app/comm-parse"
	"app/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

type ManagerController struct {
	ManagerService domain.IManagerService
	AuthService    domain.IAuthService
}

func (m *ManagerController) Create(lg *logrus.Logger) error {
	mngrDto, err := cli_ui.InputManager(lg)
	if err != nil {
		lg.Warnf("bad input manager")
		return xerrors.Errorf("manager controller: create error: %v", err.Error())
	}

	mngr := comm_parse.ToDomainManager(&mngrDto)

	_, err = m.ManagerService.Create(&mngr, lg)
	if err != nil {
		lg.Warnf("bad create manager")
		return xerrors.Errorf("manager controller: create error: %v", err.Error())
	}

	return nil
}

func (m *ManagerController) Delete(id int, lg *logrus.Logger) error {
	err := m.ManagerService.Delete(id, lg)
	if err != nil {
		lg.Warnf("bad delete manager")
		return xerrors.Errorf("manager controller: delete error: %v", err.Error())
	}
	return nil
}

func (m *ManagerController) GetByNameAndSurname(name string, surname string, lg *logrus.Logger) error {
	mngrs, err := m.ManagerService.GetByNameSurname(name, surname, lg)
	if err != nil {
		lg.Warnf("bad get by name and surname error")
		return xerrors.Errorf("manager controller: get by name and surname error: %v", err.Error())
	}
	cli_ui.OutputManagers(mngrs)
	return nil
}

func (m *ManagerController) GetByDepartment(department string, lg *logrus.Logger) error {
	mngrs, err := m.ManagerService.GetByDepartment(department, lg)
	if err != nil {
		lg.Warnf("bad get by department")
		return xerrors.Errorf("manager controller: get by department error: %v", err.Error())
	}
	cli_ui.OutputManagers(mngrs)
	return nil
}

func (m *ManagerController) GetById(id int, lg *logrus.Logger) error {
	mngr, err := m.ManagerService.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get by id manager")
		return xerrors.Errorf("manager controller: get by id error: %v", err.Error())
	}
	cli_ui.PrintManagerHeader()
	cli_ui.OutputManager(&mngr)
	return nil
}

func (m *ManagerController) GetStatistics(from time.Time, to time.Time, offset int, limit int, lg *logrus.Logger) error {
	stats, err := m.ManagerService.GetStatisticsForManagers(from, to, offset, limit, lg)
	if err != nil {
		lg.Warnf("get managers statistics error")
		return xerrors.Errorf("manager controller: get statistics error: %v", err.Error())
	}
	cli_ui.OutputStats(stats)
	return nil
}

func (m *ManagerController) ProcessManagers(choice string, lg *logrus.Logger) error {
	var err error

	accessLevels := map[string]string{
		"1": domain.AdminUser,
		"3": domain.AdminUser,
		"2": domain.AdminUser,
		"4": domain.AdminUser,
		"0": domain.GuestUser,
	}

	status, err := m.AuthService.CheckAccessRights(domain.TokenString, accessLevels[choice], lg)
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
		fmt.Print("Введите имя и фамилию менеджера: ")
		_, err = fmt.Scanf("%s %s", &name, &surname)
		if err != nil {
			err = xerrors.New("Некорректные имя/фамилия")
			break
		}
		name, surname = strings.TrimSpace(name), strings.TrimSpace(surname)
		err = m.GetByNameAndSurname(name, surname, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении менеджеров по имени и фамилии")
		}
	case "2":
		var department string
		fmt.Printf("Введите название отдела: ")
		_, err = fmt.Scanf("%s", &department)
		if err != nil {
			err = xerrors.New("Некорректные имя/фамилия")
			break
		}
		department = strings.TrimSpace(department)
		err = m.GetByDepartment(department, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении менеджеров по отделу")
		}
	case "3":
		var id int
		fmt.Print("Введите id менеджера: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = m.GetById(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении менеджера по id")
		}
	case "4":
		var fromTime, toTime time.Time
		fromTime, toTime, err := cli_ui.InputTimeline(lg)
		if err != nil {
			err = xerrors.New("Ошибка при вводе временного промежутка")
			break
		}
		var offset, limit int
		fmt.Print("Введите номер страницы: ")
		_, err = fmt.Scanf("%d", &offset)
		if err != nil {
			err = xerrors.New("Некорректный номер страницы")
			break
		}
		fmt.Print("Введите количество менеджеров на странице: ")
		_, err = fmt.Scanf("%d", &limit)
		if err != nil {
			err = xerrors.New("Некорректное количество менеджеров на странице")
			break
		}
		err = m.GetStatistics(fromTime, toTime, offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении статистики")
		}
	}
	return err
}
