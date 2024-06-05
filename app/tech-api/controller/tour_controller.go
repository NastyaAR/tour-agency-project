package controller

import (
	cli_ui "app/cli-ui"
	"app/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

type TourController struct {
	TourService    domain.ITourService
	RequestService domain.IRequestService
	AuthService    domain.IAuthService
}

func (t *TourController) Create(lg *logrus.Logger) error {
	tour, err := cli_ui.InputTour(lg)
	if err != nil {
		lg.Warnf("bad input tour")
		return xerrors.Errorf("tour controller: create error: %v", err.Error())
	}

	err = t.TourService.Create(&tour, lg)
	if err != nil {
		lg.Warnf("bad create tour")
		return xerrors.Errorf("tour controller: create error: %v", err.Error())
	}

	return nil
}

func (t *TourController) Delete(id int, lg *logrus.Logger) error {
	err := t.TourService.Delete(id, lg)
	if err != nil {
		lg.Warnf("bad delete tour")
		return xerrors.Errorf("tour controller: delete error: %v", err.Error())
	}
	return nil
}

func (t *TourController) Get(offset int, limit int, lg *logrus.Logger) error {
	tours, err := t.TourService.GetTours(offset, limit, lg)
	if err != nil {
		lg.Warnf("bad get tours")
		return xerrors.Errorf("tour controller: get error: %v", err.Error())
	}

	cli_ui.OutputTours(tours)
	return nil
}

func (t *TourController) GetFilter(offset int, limit int, lg *logrus.Logger) error {
	filter, err := cli_ui.InputFilter(lg)
	if err != nil {
		lg.Warnf("bad get tour filter")
		return xerrors.Errorf("tour controller: get filter error: %v", err.Error())
	}

	tours, err := t.TourService.GetByCriteria(&filter, offset, limit, lg)
	cli_ui.OutputTours(tours)
	return nil
}

func (t *TourController) GetHotTours(offset int, limit int, lg *logrus.Logger) error {
	tours, err := t.TourService.GetHotTours(offset, limit, lg)
	if err != nil {
		lg.Warnf("bad get hot tours")
		return xerrors.Errorf("tour controller: get hot tours error: %v", err.Error())
	}
	cli_ui.OutputHotTours(tours)
	return nil
}

func (t *TourController) SetSale(tourId int, saleId int, lg *logrus.Logger) error {
	sale := domain.Sale{
		ID:          saleId,
		Name:        "",
		ExpiredTime: time.Time{},
		Percent:     0,
	}

	err := t.TourService.SetSale(tourId, &sale, lg)
	if err != nil {
		lg.Warnf("bad set sale on tour")
		return xerrors.Errorf("tour controller: set sale error: %v", err.Error())
	}
	return nil
}

func (t *TourController) GetById(id int, lg *logrus.Logger) error {
	tour, err := t.TourService.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get tour by id")
		return xerrors.Errorf("tour controller: get by id error: %v", err.Error())
	}
	cli_ui.PrintHeader()
	cli_ui.OutputTour(&tour)
	return nil
}

func (t *TourController) CreateRequestForTour(id int, lg *logrus.Logger) error {
	clntId, err := t.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad extract clnt id from token")
		return xerrors.Errorf("tour controller: buy tour error: %v", err.Error())
	}

	req := domain.Request{
		ID:         0,
		TourID:     id,
		ClntID:     clntId,
		MngrID:     0, //менеджеров по id делает рандом
		Status:     domain.Accepted,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
		Data:       "{}",
	}

	err = t.RequestService.Create(&req, lg)
	if err != nil {
		lg.Warnf("bad create request for tour")
		return xerrors.Errorf("tour controller: buy tour error: %v", err.Error())
	}

	return nil
}

func (tc *TourController) ProcessTours(choice string, lg *logrus.Logger) error {
	var err error

	accessLevels := map[string]string{
		"1": domain.AdminUser,
		"2": domain.AdminUser,
		"3": domain.GuestUser,
		"4": domain.GuestUser,
		"5": domain.AdminUser,
		"6": domain.ManagerUser,
		"7": domain.GuestUser,
		"8": domain.ClientUser,
		"0": domain.GuestUser,
	}

	status, err := tc.AuthService.CheckAccessRights(domain.TokenString, accessLevels[choice], lg)
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
		err = tc.Create(lg)
		if err != nil {
			err = xerrors.Errorf("Некорректный id: %v", err.Error())
			break
		}
	case "2":
		var id int
		fmt.Print("Введите id удаляемой скидки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = tc.Delete(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при удалении скидки")
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
		err = tc.Get(offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при выводе скидок")
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
		err = tc.GetFilter(offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при выводе скидок")
		}
	case "5":
		var id int
		fmt.Print("Введите id скидки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id скидки")
			break
		}
		var saleId int
		fmt.Print("Введите id скидки: ")
		_, err = fmt.Scanf("%d", &saleId)
		if err != nil {
			err = xerrors.New("Некорректный id скидки")
			break
		}
		err = tc.SetSale(id, saleId, lg)
		if err != nil {
			err = xerrors.New("Ошибка при добавлении скидки")
		}
	case "6":
		var id int
		fmt.Print("Введите id скидки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = tc.GetById(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при выводе скидки по id")
		}
	case "7":
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
		err = tc.GetHotTours(offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении горячих предложений")
		}
	case "8":
		var id int
		fmt.Print("Введите id тура: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil || id <= 0 {
			err = xerrors.New("Некорректный id")
			break
		}
		err = tc.CreateRequestForTour(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при создании заявки на тур")
		}
	case "0":
		err = nil
	}

	return err
}
