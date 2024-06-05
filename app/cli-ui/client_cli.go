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
)

func InputClient(lg *logrus.Logger) (domain.NewAccountDTO, error) {
	var name, surname, mail, phone string
	var err error

	fmt.Println("\nВведите данные клиента:\n")
	fmt.Print("Имя: ")
	name, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input name client: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input client error: %v", err.Error())
	}
	name = strings.TrimSpace(name)
	fmt.Print("Фамилия: ")
	surname, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input surname client: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input client error: %v", err.Error())
	}
	surname = strings.TrimSpace(surname)
	fmt.Print("Почта: ")
	mail, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input mail client: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input client error: %v", err.Error())
	}
	mail = strings.TrimSpace(mail)
	fmt.Print("Номер телефона: ")
	phone, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input phone client: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input client error: %v", err.Error())
	}
	phone = strings.TrimSpace(phone)

	ParseNewAccString := fmt.Sprintf(comm_parse.ParseClientString, 0, name, surname, mail, phone)
	clntJson, err := comm_parse.FromStringToUserJson(ParseNewAccString, lg)
	if err != nil {
		lg.Warnf("bad input client: %v", err.Error())
		return domain.NewAccountDTO{}, xerrors.Errorf("input client error: %v", err.Error())
	}

	clnt := clntJson.ToDtoClientAccount(lg)
	return clnt, nil
}

func PrintClientHeader() {
	fmt.Printf("\n%-5s %-25s %-25s %-25s %-25s\n", "ID", "Имя",
		"Фамилия", "Почта", "Телефон")
}

func OutputClient(s *domain.Client) {
	fmt.Printf("%-5d %-25s %-25s %-25s %-25s\n", s.ID, s.Name,
		s.Surname, s.Mail, s.Phone)
}

func OutputClients(clnts []domain.Client) {
	PrintClientHeader()
	fmt.Printf("%s\n", strings.Repeat("-", 109))
	for _, clnt := range clnts {
		OutputClient(&clnt)
	}
	fmt.Print("\n")
}

func GetClientChoice() string {
	fmt.Println("\n\nВыберите один из пунктов меню:\n")
	fmt.Println("1. отфильтровать клиентов по имени и фамилии")
	fmt.Println("2. отфильтровать клиентов по номеру телефона")
	fmt.Println("3. получить клиента по id")
	fmt.Println("4. получить историю заявок")
	fmt.Println("0. вернуться\n")

	fmt.Print("\nВведите номер пункта: ")
	choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(choice)
}
