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

func GetSaleChoice() string {
	fmt.Println("\n\nВыберите один из пунктов меню:\n")
	fmt.Println("1. создать скидочное предложение")
	fmt.Println("2. удалить скидочное предложение")
	fmt.Println("3. вывести скидочные предложения")
	fmt.Println("4. отфильтровать скидочные предложения")
	fmt.Println("5. получить скидочное предложение по id")
	fmt.Println("0. вернуться\n")

	fmt.Print("\nВведите номер пункта: ")
	choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(choice)
}

func InputSale(lg *logrus.Logger) (domain.Sale, error) {
	var name, expiredTime string
	var percent int
	var err error

	fmt.Println("\nВведите данные для создания скидочного предложения")
	fmt.Print("Название скидки: ")
	name, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad sale input: %v", err.Error())
		return domain.Sale{}, xerrors.Errorf("input sale error: %v", err.Error())
	}
	name = strings.TrimSpace(name)
	fmt.Print("Дата истечения: ")
	expiredTime, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input sale: %v", err.Error())
		return domain.Sale{}, xerrors.Errorf("input sale error: %v", err.Error())
	}
	expiredTime = strings.TrimSpace(expiredTime)
	fmt.Print("Процент скидки: ")
	_, err = fmt.Scanln(&percent)
	if err != nil {
		lg.Warnf("bad input tour: %v", err.Error())
		return domain.Sale{}, xerrors.Errorf("input sale error: %v", err.Error())
	}

	ParseSaleString := fmt.Sprintf(comm_parse.ParseSaleString, 0, name, expiredTime, percent)

	saleJson, err := comm_parse.FromStringToSaleJson(ParseSaleString, lg)
	if err != nil {
		lg.Warnf("bad input sale: %v", err.Error())
		return domain.Sale{}, xerrors.Errorf("input sale error: %v", err.Error())
	}

	sale, err := saleJson.ToDomainSale(lg)
	if err != nil {
		lg.Warnf("bad input sale: %v", err.Error())
		return domain.Sale{}, xerrors.Errorf("input sale error: %v", err.Error())
	}

	return sale, nil
}

func PrintSaleHeader() {
	fmt.Printf("\n%-5s %-25s %-25s %-20s\n", "ID", "Название",
		"Дата истечения", "Процент скидки")
}

func OutputSale(s *domain.Sale) {
	fmt.Printf("%-5d %-25s %-25s %-20d\n", s.ID, s.Name,
		s.ExpiredTime.Format("2006-01-02 15:04"), s.Percent)
}

func OutputSales(sales []domain.Sale) {
	PrintSaleHeader()
	fmt.Printf("%s\n", strings.Repeat("-", 78))
	for _, sale := range sales {
		OutputSale(&sale)
	}
	fmt.Print("\n")
}

func InputSaleFilter(lg *logrus.Logger) (domain.Sale, error) {
	sale := domain.Sale{
		ID:   0,
		Name: "",
		ExpiredTime: time.Date(1970, 1,
			1, 0, 0, 0, 0, time.UTC),
		Percent: domain.DefaultEmptyValue,
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
			return domain.Sale{}, xerrors.Errorf("input filter error: %v", err.Error())
		}
		switch n {
		case "1":
			sale.Name = text
		case "2":
			layout := "2006-01-02 15:04"
			timeLocal, err := time.Parse(layout, text)
			if err != nil {
				lg.Warnf("bad sale filter date")
				return domain.Sale{}, xerrors.Errorf("sale filter error: %v", err.Error())
			}
			sale.ExpiredTime = timeLocal
		case "3":
			sale.Percent, err = strconv.Atoi(text)
			if err != nil {
				lg.Warnf("bad sale duration")
				return domain.Sale{}, xerrors.Errorf("sale filter error: %v", err.Error())
			}
		}
	}

	return sale, nil
}
