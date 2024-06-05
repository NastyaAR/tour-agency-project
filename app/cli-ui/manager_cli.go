package cli_ui

import (
	comm_parse "app/comm-parse"
	"app/domain"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"os"
	"strings"
	"time"
)

func InputManager(lg *logrus.Logger) (domain.NewAccountDTO, error) {
	var name, surname, department string
	var err error

	fmt.Println("\nВведите данные менеджера:\n")
	fmt.Print("Имя: ")
	name, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input name manager: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input manager error: %v", err.Error())
	}
	name = strings.TrimSpace(name)
	fmt.Print("Фамилия: ")
	surname, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input surname manager: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input manager error: %v", err.Error())
	}
	surname = strings.TrimSpace(surname)
	fmt.Print("Подразделение: ")
	department, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input department manager: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input manager error: %v", err.Error())
	}
	department = strings.TrimSpace(department)

	ParseNewAccString := fmt.Sprintf(comm_parse.ParseManagerString, 0, name, surname, department)
	mngrJson, err := comm_parse.FromStringToUserJson(ParseNewAccString, lg)
	if err != nil {
		lg.Warnf("bad input manager: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input manager error: %v", err.Error())
	}

	mngr := mngrJson.ToDtoManagerAccount(lg)
	return mngr, nil
}

func PrintManagerHeader() {
	fmt.Printf("\n%-5s %-25s %-25s %-25s\n", "ID", "Имя",
		"Фамилия", "Подразделение")
}

func OutputManager(s *domain.Manager) {
	fmt.Printf("%-5d %-25s %-25s %-25s\n", s.ID, s.Name,
		s.Surname, s.Department)
}

func OutputManagers(mngrs []domain.Manager) {
	PrintClientHeader()
	fmt.Printf("%s\n", strings.Repeat("-", 109))
	for _, mngr := range mngrs {
		OutputManager(&mngr)
	}
	fmt.Print("\n")
}

func PrintStatisticsHeader() {
	fmt.Printf("\n%-5s %-25s %-25s %-25s %-15s %-15s\n", "Номер",
		"Фамилия", "Имя", "Подразделение", "Эффективность", "Сумма продаж")
}

func OutputStat(n int, st *domain.Statistics) {
	fmt.Printf("%-5d, %-25s %-25s %-25s %-15f %-15d\n", n, st.ManagerAcc.Surname,
		st.ManagerAcc.Name, st.ManagerAcc.Department, st.Efficiency, st.SaleSum)
}

func OutputStats(stats []domain.Statistics) {
	PrintStatisticsHeader()
	fmt.Printf("%s\n", strings.Repeat("-", 115))
	for i, stat := range stats {
		OutputStat(i, &stat)
	}
	fmt.Print("\n")
}

func GetManagerChoice() string {
	fmt.Println("\n\nВыберите один из пунктов меню:\n")
	fmt.Println("1. отфильтровать менеджеров по имени и фамилии")
	fmt.Println("2. отфильтровать менеджеров по отделу")
	fmt.Println("3. получить менеджера по id")
	fmt.Println("4. получить статистику по работе менеджеров за период")
	fmt.Println("0. вернуться\n")

	fmt.Print("\nВведите номер пункта: ")
	choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(choice)
}

func InputPartOfTimeLine(prompt string, lg *logrus.Logger) (time.Time, error) {
	var timeStr string
	fmt.Printf(prompt)
	timeStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	timeStr = strings.TrimSpace(timeStr)
	if err != nil {
		lg.Warnf("input timeline error")
		return time.Date(1970, 1,
				1, 0, 0, 0, 0, time.UTC),
			xerrors.Errorf("input timeline error: %v", err.Error())
	}

	layout := "2006-01-02"
	timeLocal, err := time.Parse(layout, timeStr)
	if err != nil {
		lg.Warnf("input timeline error")
		return time.Date(1970, 1,
				1, 0, 0, 0, 0, time.UTC),
			xerrors.Errorf("input timeline error: %v", err.Error())
	}
	return timeLocal, nil
}

func InputTimeline(lg *logrus.Logger) (time.Time, time.Time, error) {
	fmt.Printf("Введите время в формате ГГГГ-ММ-ДД: \n")
	fromStr, err := InputPartOfTimeLine("От: ", lg)
	if err != nil {
		lg.Warnf("input part of timeline error: %v", err.Error())
		return time.Date(1970, 1,
				1, 0, 0, 0, 0, time.UTC), time.Date(1970, 1,
				1, 0, 0, 0, 0, time.UTC),
			xerrors.Errorf("input timeline error: %v", err.Error())
	}

	toStr, err := InputPartOfTimeLine("От: ", lg)
	if err != nil {
		lg.Warnf("input part of timeline error: %v", err.Error())
		return time.Date(1970, 1,
				1, 0, 0, 0, 0, time.UTC), time.Date(1970, 1,
				1, 0, 0, 0, 0, time.UTC),
			xerrors.Errorf("input timeline error: %v", err.Error())
	}

	return fromStr, toStr, nil
}
