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

func GetAccountChoice() string {
	fmt.Println("\n\nВыберите один из пунктов меню:\n")
	fmt.Println("1. зарегистрироваться")
	fmt.Println("2. зарегистрировать менеджера")
	fmt.Println("3. войти как клиент")
	fmt.Println("4. войти как менеджер")
	fmt.Println("5. войти как администратор")
	fmt.Println("6. выйти")
	fmt.Println("0. вернуться\n")

	fmt.Print("\nВведите номер пункта: ")
	choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(choice)
}

func InputAccountForLogin(lg *logrus.Logger) (domain.Account, error) {
	var login, password string
	var err error

	fmt.Println("\nВведите данные для входа\n")
	fmt.Print("Логин: ")
	login, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input login: %v", err.Error())
		return domain.Account{}, xerrors.Errorf("input account error: %v", err.Error())
	}
	login = strings.TrimSpace(login)
	fmt.Print("Пароль: ")
	password, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input password: %v", err.Error())
		return domain.Account{}, xerrors.Errorf("input account error: %v", err.Error())
	}
	password = strings.TrimSpace(password)

	ParseAccString := fmt.Sprintf(comm_parse.ParseAccountString, 0, login, password)
	accJson, err := comm_parse.FromStringToAccountJson(ParseAccString, lg)
	if err != nil {
		lg.Warnf("bad input account: %v", err.Error())
		return domain.Account{}, xerrors.Errorf("input account error: %v", err.Error())
	}

	acc := accJson.ToDomainAccount(lg)
	return acc, nil
}

func InputAccount(lg *logrus.Logger) (domain.Account, error) {
	var login, password, passAgain string
	var err error

	fmt.Println("\nВведите данные для создания аккаунта\n")
	fmt.Print("Логин: ")
	login, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input login: %v", err.Error())
		return domain.Account{}, xerrors.Errorf("input account error: %v", err.Error())
	}
	login = strings.TrimSpace(login)
	fmt.Print("Пароль: ")
	password, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input password: %v", err.Error())
		return domain.Account{}, xerrors.Errorf("input account error: %v", err.Error())
	}
	password = strings.TrimSpace(password)
	fmt.Print("Введите пароль ещё раз: ")
	passAgain, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		lg.Warnf("bad input password again: %v", err.Error())
		return domain.Account{}, xerrors.Errorf("input account error: %v", err.Error())
	}
	passAgain = strings.TrimSpace(passAgain)

	if password != passAgain {
		return domain.Account{}, xerrors.Errorf("input account error: password doest same")
	}

	ParseAccString := fmt.Sprintf(comm_parse.ParseAccountString, 0, login, password)
	accJson, err := comm_parse.FromStringToAccountJson(ParseAccString, lg)
	if err != nil {
		lg.Warnf("bad input account: %v", err.Error())
		return domain.Account{}, xerrors.Errorf("input account error: %v", err.Error())
	}

	acc := accJson.ToDomainAccount(lg)
	return acc, nil
}
