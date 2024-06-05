package cli_ui

import (
	comm_parse "app/comm-parse"
	"app/domain"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"os"
	"strconv"
	"strings"
)

func GetRequestChoice() string {
	fmt.Println("\n\nВыберите один из пунктов меню:\n")
	fmt.Println("1. создать заявку на подбор тура")
	fmt.Println("2. вывести данные заявки по id")
	fmt.Println("3. добавить тур в заявку")
	fmt.Println("4. получить заявку по id")
	fmt.Println("5. получить заявки по статусу")
	fmt.Println("6. получить заявки")
	fmt.Println("7. подтвердить заявку")
	fmt.Println("8. отклонить заявку")
	fmt.Println("9. оплатить заявку")
	fmt.Println("10. получить мои заявки")
	fmt.Println("0. вернуться\n")

	fmt.Print("\nВведите номер пункта: ")
	choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(choice)
}

func InputRequest(clntId int, lg *logrus.Logger) (domain.Request, error) {
	var data string
	var err error

	fmt.Println("\nВведите данные для создания заявки")

	tour, err := InputRequestData(lg)
	if err != nil {
		lg.Warnf("bad input request")
		return domain.Request{}, xerrors.Errorf("request cli: input request error: %v", err.Error())
	}

	data = comm_parse.FromDomainTourToJsonString(tour)

	var req domain.Request
	req.Data = data
	req.ClntID = clntId
	req.TourID = domain.DefaultEmptyValue
	return req, nil
}

func PrintRequestHeader() {
	fmt.Printf("\n%-5s %-10s %-10s %-10s %-15s %-35s %-35s\n",
		"ID", "ID тура", "ID клиента", "ID менеджера",
		"Статус", "Время создания", "Время изменения")
}

func OutputRequest(req *domain.Request) {
	fmt.Printf("%-5d %-10d %-10d %-10d %-15s %-35s %-35s\n",
		req.ID, req.TourID, req.ClntID, req.MngrID, req.Status, req.CreateTime,
		req.ModifyTime)
	if req.Data != "{}" {
		var result map[string]interface{}
		_ = json.Unmarshal([]byte(req.Data), &result)
		fmt.Println()
		for k, v := range result {
			fmt.Println(k, ": ", v)
		}
	}
}

func OutputRequestData(req *domain.Request) {
	fmt.Printf(req.Data)
}

func OutputRequests(reqs []domain.Request) {
	PrintRequestHeader()
	fmt.Printf("%s\n", strings.Repeat("-", 127))
	for _, req := range reqs {
		OutputRequest(&req)
	}
	fmt.Print("\n")
}

func InputRequestData(lg *logrus.Logger) (domain.Tour, error) {
	var duration, cost, touristsNumber int
	var chillPlace, fromPlace, date, chillType, durationStr, costStr, touristsNumberStr string
	var err error

	fmt.Println("\nВведите данные для создания заявки:")
	fmt.Print("Место отдыха: ")
	chillPlace, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	chillPlace = strings.TrimSpace(chillPlace)

	fmt.Print("Откуда: ")
	fromPlace, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	fromPlace = strings.TrimSpace(fromPlace)
	fmt.Print("Дата: ")
	date, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	date = strings.TrimSpace(date)
	fmt.Print("Продолжительность: ")
	durationStr, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	durationStr = strings.TrimSpace(durationStr)
	if durationStr != "" {
		duration, err = strconv.Atoi(strings.TrimSpace(durationStr))
		if err != nil {
			lg.Warnf("bad input tour: %v", err.Error())
			return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
		}
	} else {
		duration = domain.DefaultEmptyValue
	}
	fmt.Print("Стоимость: ")
	costStr, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	costStr = strings.TrimSpace(costStr)
	if costStr != "" {
		cost, err = strconv.Atoi(strings.TrimSpace(costStr))
		if err != nil {
			lg.Warnf("bad input tour: %v", err.Error())
			return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
		}
	} else {
		cost = domain.DefaultEmptyValue
	}
	fmt.Print("Количество туристов: ")
	touristsNumberStr, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	touristsNumberStr = strings.TrimSpace(touristsNumberStr)
	if touristsNumberStr != "" {
		touristsNumber, err = strconv.Atoi(strings.TrimSpace(touristsNumberStr))
		if err != nil {
			lg.Warnf("bad input tour: %v", err.Error())
			return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
		}
	} else {
		touristsNumber = domain.DefaultEmptyValue
	}
	fmt.Print("Тип отдыха: ")
	chillType, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	chillType = strings.TrimSpace(chillType)

	ParseTourString := fmt.Sprintf(comm_parse.ParseTourString, 0, chillPlace, fromPlace, date,
		duration, cost, touristsNumber, chillType)

	tourJson, err := comm_parse.FromStringToTourJson(ParseTourString, lg)
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}

	tour, err := tourJson.ToDomainTour(lg)
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}

	return tour, nil
}
