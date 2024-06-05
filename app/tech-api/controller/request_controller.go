package controller

import (
	cli_ui "app/cli-ui"
	"app/domain"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"os"
	"strings"
)

type RequestController struct {
	RequestService domain.IRequestService
	AuthService    domain.IAuthService
}

func (r *RequestController) Create(lg *logrus.Logger) error {
	id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad extract role from token")
		return xerrors.Errorf("request controller: create error: %v", err.Error())
	}

	req, err := cli_ui.InputRequest(id, lg)
	if err != nil {
		lg.Warnf("bad input request")
		return xerrors.Errorf("request controller: create error: %v", err.Error())
	}

	err = r.RequestService.Create(&req, lg)
	if err != nil {
		lg.Warnf("bad create request")
		return xerrors.Errorf("request controller: create error: %v", err.Error())
	}

	return nil
}

func (r *RequestController) GetById(id int, lg *logrus.Logger) error {
	req, err := r.RequestService.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get by id request")
		return xerrors.Errorf("request controller: get by id error: %v", err.Error())
	}
	cli_ui.PrintRequestHeader()
	cli_ui.OutputRequest(&req)
	return nil
}

func (r *RequestController) GetInfoById(id int, lg *logrus.Logger) error {
	req, err := r.RequestService.GetById(id, lg)
	if err != nil {
		lg.Warnf("bad get by id request")
		return xerrors.Errorf("request controller: get by id error: %v", err.Error())
	}
	cli_ui.OutputRequestData(&req)
	return nil
}

func (r *RequestController) AddTour(reqId int, tourId int, lg *logrus.Logger) error {
	id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad add tour: %v", err.Error())
		return xerrors.Errorf("request controller: add tour error: %v", err.Error())
	}
	err = r.RequestService.AddTour(reqId, tourId, id, lg)
	if err != nil {
		lg.Warnf("bad update request")
		return xerrors.Errorf("request controller: add tour error: %v", err.Error())
	}
	return nil
}

func (r *RequestController) GetByStatus(status string, offset int, limit int, lg *logrus.Logger) error {
	lg.Info("request controller: get by status")
	mngr_id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("request controller: get by status warn: bad get by status requests: %v", err.Error())
		return xerrors.Errorf("request controller: get by status error: %v", err.Error())
	}
	reqs, err := r.RequestService.GetByStatus(mngr_id, status, offset, limit, lg)
	if err != nil {
		lg.Errorf("request controller: get by status error: bad get by status request: %v", err.Error())
		return xerrors.Errorf("request controller: get by status error: %v", err.Error())
	}
	cli_ui.OutputRequests(reqs)
	return nil
}

func (r *RequestController) Get(offset int, limit int, lg *logrus.Logger) error {
	lg.Info("request controller: get")
	mngr_id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("request controller: get warn: bad get requests: %v", err.Error())
		return xerrors.Errorf("request controller: get error: %v", err.Error())
	}
	reqs, err := r.RequestService.GetRequests(mngr_id, offset, limit, lg)
	if err != nil {
		lg.Warnf("request controller: get error: bad get requests: %v", err.Error())
		return xerrors.Errorf("request controller: get error: %v", err.Error())
	}
	cli_ui.OutputRequests(reqs)
	return nil
}

func (r *RequestController) Reject(id int, lg *logrus.Logger) error {
	mngr_id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad reject tour: %v", err.Error())
		return xerrors.Errorf("request controller: reject error: %v", err.Error())
	}
	err = r.RequestService.Reject(id, mngr_id, lg)
	if err != nil {
		lg.Warnf("bad reject request")
		return xerrors.Errorf("request controller: reject error: %v", err.Error())
	}
	return nil
}

func (r *RequestController) Approve(id int, lg *logrus.Logger) error {
	mngr_id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad approve tour: %v", err.Error())
		return xerrors.Errorf("request controller: approve error: %v", err.Error())
	}
	err = r.RequestService.Approve(id, mngr_id, lg)
	if err != nil {
		lg.Warnf("bad approve request")
		return xerrors.Errorf("request controller: approve error: %v", err.Error())
	}
	return nil
}

func (r *RequestController) Pay(id int, lg *logrus.Logger) error {
	clnt_id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad pay tour: %v", err.Error())
		return xerrors.Errorf("request controller: pay error: %v", err.Error())
	}

	err = r.RequestService.Pay(id, clnt_id, lg)
	if err != nil {
		lg.Warnf("bad pay request")
		return xerrors.Errorf("request controller: pay error: %v", err.Error())
	}
	return nil
}

func (r *RequestController) GetRequestsForClient(offset int, limit int, lg *logrus.Logger) error {
	id, err := r.AuthService.ExtractIdFromToken(domain.TokenString, lg)
	if err != nil {
		lg.Warnf("bad get requests for client: %v", err.Error())
		return xerrors.Errorf("request controller: get requests for client error: %v", err.Error())
	}
	fmt.Println(id)
	reqs, err := r.RequestService.GetRequestsForClient(id, offset, limit, lg)
	if err != nil {
		lg.Warnf("bad get requests for client: %v", err.Error())
		return xerrors.Errorf("request controller: get requests for client error: %v", err.Error())
	}
	cli_ui.OutputRequests(reqs)
	return nil
}

func (r *RequestController) ProcessRequests(choice string, lg *logrus.Logger) error {
	var err error
	var reqStatus string

	accessLevels := map[string]string{
		"1":  domain.ClientUser,
		"2":  domain.ManagerUser,
		"3":  domain.ManagerUser,
		"4":  domain.ManagerUser,
		"5":  domain.ManagerUser,
		"6":  domain.ManagerUser,
		"7":  domain.ManagerUser,
		"8":  domain.ManagerUser,
		"9":  domain.ClientUser,
		"10": domain.ClientUser,
		"0":  domain.GuestUser,
	}

	status, err := r.AuthService.CheckAccessRights(domain.TokenString, accessLevels[choice], lg)
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
		err = r.Create(lg)
		if err != nil {
			err = xerrors.New("Ошибка при создании заявки")
		}
	case "2":
		var id int
		fmt.Print("Введите id заявки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = r.GetInfoById(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении данных заявки")
		}
	case "3":
		var id, tourId int
		fmt.Print("Введите id заявки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		fmt.Print("Введите id тура: ")
		_, err = fmt.Scanf("%d", &tourId)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = r.AddTour(id, tourId, lg)
		if err != nil {
			err = xerrors.New("Ошибка при добавлении тура в заявку")
		}
	case "4":
		var id int
		fmt.Print("Введите id заявки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = r.GetById(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении заявки по id")
		}
	case "5":
		reqStatus, err = bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			err = xerrors.New("Некорректный статус заявки")
			break
		}
		var offset, limit int
		fmt.Print("Введите номер страницы: ")
		_, err = fmt.Scanf("%d", &offset)
		if err != nil {
			err = xerrors.New("Некорректный номер страницы")
			break
		}
		fmt.Print("Введите количество заявок на странице: ")
		_, err = fmt.Scanf("%d", &limit)
		if err != nil {
			err = xerrors.New("Некорректное количество заявок на странице")
			break
		}
		reqStatus = strings.TrimSpace(reqStatus)
		err = r.GetByStatus(reqStatus, offset-1, limit, lg)
		if err != nil {
			err = xerrors.Errorf("Ошибка вывода заявок по статусу: %v", err.Error())
		}
	case "6":
		var offset, limit int
		fmt.Print("Введите номер страницы: ")
		_, err = fmt.Scanf("%d", &offset)
		if err != nil {
			err = xerrors.New("Некорректный номер страницы")
			break
		}
		fmt.Print("Введите количество заявок на странице: ")
		_, err = fmt.Scanf("%d", &limit)
		if err != nil {
			err = xerrors.New("Некорректное количество заявок на странице")
			break
		}
		err = r.Get(offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при получении заявок")
		}
	case "7":
		var id int
		fmt.Print("Введите id заявки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = r.Approve(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при подтверждении заявки")
		}
	case "8":
		var id int
		fmt.Print("Введите id заявки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = r.Reject(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при отклонении заявки")
		}
	case "9":
		var id int
		fmt.Print("Введите id заявки: ")
		_, err = fmt.Scanf("%d", &id)
		if err != nil {
			err = xerrors.New("Некорректный id")
			break
		}
		err = r.Pay(id, lg)
		if err != nil {
			err = xerrors.New("Ошибка при оплате заявки")
		} else {
			fmt.Println("\nОплата тура прошла успешно!")
		}
	case "10":
		var offset, limit int
		fmt.Print("Введите номер страницы: ")
		_, err = fmt.Scanf("%d", &offset)
		if err != nil {
			err = xerrors.New("Некорректный номер страницы")
			break
		}
		fmt.Print("Введите количество заявок на странице: ")
		_, err = fmt.Scanf("%d", &limit)
		if err != nil {
			err = xerrors.New("Некорректное количество заявок на странице")
			break
		}
		err = r.GetRequestsForClient(offset-1, limit, lg)
		if err != nil {
			err = xerrors.New("Ошибка при выводе заявок клиента")
		}
	}

	return err
}
