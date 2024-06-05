package controller

import (
	cli_ui "app/cli-ui"
	"app/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
)

type SaleController struct {
	SaleService domain.ISaleService
	AuthService domain.IAuthService
}

func (s *SaleController) Create(lg *logrus.Logger) error {
	sale, err := cli_ui.InputSale(lg)
	if err != nil {
		lg.Warnf("bad input sale")
		return xerrors.Errorf("sale controller: create error: %v", err.Error())
	}

	err = s.SaleService.Create(&sale, lg)
	if err != nil {
		lg.Warnf("bad create sale")
		return xerrors.Errorf("sale controller: create error: %v", err.Error())
	}

	return nil
}

func (s *SaleController) Delete(id int, lg *logrus.Logger) error {
	err := s.SaleService.Delete(id, lg)
	if err != nil {
		lg.Warnf("bad delete sale")
		return xerrors.Errorf("sale controller: delete error: %v", err.Error())
	}
	return nil
}

func (s *SaleController) Get(offset int, limit int, lg *logrus.Logger) error {
	sales, err := s.SaleService.GetSales(offset, limit, lg)
	if err != nil {
		lg.Warnf("bad get sales")
		return xerrors.Errorf("sale controller: get error: %v", err.Error())
	}

	cli_ui.OutputSales(sales)
	return nil
}

func (s *SaleController) GetFilter(offset int, limit int, lg *logrus.Logger) error {
	filter, err := cli_ui.InputSaleFilter(lg)
	fmt.Println(filter)
	if err != nil {
		lg.Warnf("bad get sale filter")
		return xerrors.Errorf("sale controller: get filter error: %v", err.Error())
	}

	sales, err := s.SaleService.GetByCriteria(&filter, offset, limit, lg)
	if err != nil {
		lg.Warnf("bad get by criteria sales")
		return xerrors.Errorf("sale controller: get filter error: %v", err.Error())
	}
	cli_ui.OutputSales(sales)
	return nil
}

func (s *SaleController) GetById(id int, lg *logrus.Logger) error {
	sale, err := s.SaleService.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get sale by id")
		return xerrors.Errorf("sale controller: get by id error: %v", err.Error())
	}
	cli_ui.PrintSaleHeader()
	cli_ui.OutputSale(&sale)
	return nil
}

func (s *SaleController) ProcessSales(choice string, lg *logrus.Logger) error {
	var err error

	accessLevels := map[string]string{
		"1": domain.AdminUser,
		"2": domain.AdminUser,
		"3": domain.GuestUser,
		"4": domain.GuestUser,
		"5": domain.ManagerUser,
		"0": domain.GuestUser,
	}

	status, err := s.AuthService.CheckAccessRights(domain.TokenString, accessLevels[choice], lg)
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
		err = s.Create(lg)
		if err != nil {
			err = xerrors.New("Ошибка при создании тура")
			break
		}
	case "2":
		var id int
		fmt.Print("Введите id удаляемого скидочного предложения: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = s.Delete(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при удалении тура")
			break
		}
	case "3":
		var offset, limit int
		fmt.Print("Введите номер страницы: ")
		_, err = fmt.Scanf("%d", &offset)
		if err != nil {
			err = xerrors.New("Некорректный номер страницы")
			break
		}
		fmt.Print("Введите количество скидок на странице: ")
		_, err = fmt.Scanf("%d", &limit)
		if err != nil {
			err = xerrors.New("Некорректное количество скидок на странице")
			break
		}
		err = s.Get(offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении скидок")
			break
		}
	case "4":
		var offset, limit int
		fmt.Print("Введите номер страницы: ")
		_, err = fmt.Scanf("%d", &offset)
		if err != nil {
			err = xerrors.New("Некорректный номер страницы")
			break
		}
		fmt.Print("Введите количество скидок на странице: ")
		_, err = fmt.Scanf("%d", &limit)
		if err != nil {
			err = xerrors.New("Некорректное количество скидок на странице")
			break
		}
		err = s.GetFilter(offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при фильтрации туров")
			break
		}
	case "5":
		var id int
		fmt.Print("Введите id скидки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = s.GetById(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении скидок по id")
			break
		}
	case "0":
		err = nil
	}

	return err
}
