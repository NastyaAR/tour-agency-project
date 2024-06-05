package cli_ui

import (
	comm_parse "app/comm-parse"
	"app/domain"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"os"
	"strconv"
	"strings"
	"time"
)

func inputField(prompt string, isString bool, lg *logrus.Logger) (interface{}, error) {
	var dataInt int
	var data string
	var err error

	fmt.Print(prompt)
	if isString {
		data, err = bufio.NewReader(os.Stdin).ReadString('\n')
		data = strings.TrimSpace(data)
	} else {
		_, err = fmt.Scanln(&dataInt)
	}
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return nil, xerrors.Errorf("input tour error: %v", err.Error())
	}

	if isString {
		return data, nil
	}

	return dataInt, nil
}

func InputTour(lg *logrus.Logger) (domain.Tour, error) {
	var duration, cost, touristsNumber int
	var chillPlace, fromPlace, date, chillType string
	var err error

	fmt.Println("\nВведите данные для создания тура:")
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
	_, err = fmt.Scanln(&duration)
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	fmt.Print("Стоимость: ")
	_, err = fmt.Scanln(&cost)
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
	}
	fmt.Print("Количество туристов: ")
	_, err = fmt.Scanln(&touristsNumber)
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Tour{}, xerrors.Errorf("input tour error: %v", err.Error())
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

func PrintHeader() {
	fmt.Printf("\n%-5s %-25s %-25s %-20s %-20s %-15s %-20s %-10s\n",
		"ID", "Место отдыха", "Место отправления", "Дата", "Продолжительность", "Стоимость", "Кол-во туристов", "Тип отдыха")
}

func PrintHotHeader() {
	fmt.Printf("\n%-5s %-25s %-25s %-20s %-20s %-15s %-20s %-10s %-10s %-10s\n",
		"ID", "Место отдыха", "Место отправления", "Дата", "Продолжительность",
		"Стоимость", "Кол-во туристов", "Тип отдыха",
		"Скидка", "Процент")
}

func OutputTour(t *domain.Tour) {
	fmt.Printf("%-5d %-25s %-25s %-20s %-20d %-15d %-20d %-10s\n",
		t.ID, t.ChillPlace, t.FromPlace, t.Date.Format("2006-01-02 15:04"), t.Duration, t.Cost, t.TouristsNumber, t.ChillType)
}

func OutputHotTour(t *domain.HotTourDto) {
	fmt.Printf("%-5d %-25s %-25s %-20s %-20d %-15d %-20d %-10s %-10s %-10d\n",
		t.ID, t.ChillPlace, t.FromPlace, t.Date.Format("2006-01-02 15:04"), t.Duration, t.Cost,
		t.TouristsNumber, t.ChillType, t.SaleName, t.SalePercemt)
}

func OutputTours(tours []domain.Tour) {
	PrintHeader()
	fmt.Printf("%s\n", strings.Repeat("-", 147))
	for _, tour := range tours {
		OutputTour(&tour)
	}
	fmt.Print("\n")
}

func OutputHotTours(tours []domain.HotTourDto) {
	PrintHotHeader()
	fmt.Printf("%s\n", strings.Repeat("-", 169))
	for _, tour := range tours {
		OutputHotTour(&tour)
	}
	fmt.Print("\n")
}

func InputFilter(lg *logrus.Logger) (domain.Tour, error) {
	tour := domain.Tour{
		ID:         0,
		ChillPlace: "",
		FromPlace:  "",
		Date: time.Date(1970, 1,
			1, 0, 0, 0, 0, time.UTC),
		Duration:       domain.DefaultEmptyValue,
		Cost:           domain.DefaultEmptyValue,
		TouristsNumber: domain.DefaultEmptyValue,
		ChillType:      "",
	}

	for {
		fmt.Print("Введите номер фильтра: ")
		n, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		n = strings.TrimSpace(n)
		if n == "" {
			break
		}
		fmt.Print("Введите значение для поиска: ")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		text = strings.TrimSpace(text)
		if err != nil {
			lg.Warnf("bad input tour filter: %v", err.Error())
			return domain.Tour{}, xerrors.Errorf("input filter error: %v", err.Error())
		}
		switch n {
		case "1":
			tour.ChillPlace = text
		case "2":
			tour.FromPlace = text
		case "3":
			layout := "2006-01-02 15:04"
			timeLocal, err := time.Parse(layout, text)
			if err != nil {
				lg.Warnf("bad tour filter date")
				return domain.Tour{}, xerrors.Errorf("tour filter error: %v", err.Error())
			}
			tour.Date = timeLocal
		case "4":
			tour.Duration, err = strconv.Atoi(text)
			if err != nil {
				lg.Warnf("bad tour duration")
				return domain.Tour{}, xerrors.Errorf("tour filter error: %v", err.Error())
			}
		case "5":
			tour.Cost, err = strconv.Atoi(text)
			if err != nil {
				lg.Warnf("bad tour cost")
				return domain.Tour{}, xerrors.Errorf("tour filter error: %v", err.Error())
			}
		case "6":
			tour.TouristsNumber, err = strconv.Atoi(text)
			if err != nil {
				lg.Warnf("bad tour tourists number")
				return domain.Tour{}, xerrors.Errorf("tour filter error: %v", err.Error())
			}
		case "7":
			tour.ChillType = text
		}
	}

	return tour, nil
}

func GetTourChoice() string {
	fmt.Println("\n\nВыберите один из пунктов меню:\n")
	fmt.Println("1. создать тур")
	fmt.Println("2. удалить тур")
	fmt.Println("3. вывести туры")
	fmt.Println("4. отфильтровать туры")
	fmt.Println("5. поставить скидку")
	fmt.Println("6. получить тур по id")
	fmt.Println("7. вывести горячие предложения")
	fmt.Println("8. купить тур")
	fmt.Println("0. вернуться\n")

	fmt.Print("\nВведите номер пункта: ")
	choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(choice)
}
