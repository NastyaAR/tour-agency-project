package cli_ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetSectionChoice() string {
	fmt.Println("\n\nВыберите раздел:")
	fmt.Println("1. Туры")
	fmt.Println("2. Горячие предложения")
	fmt.Println("3. Заявки")
	fmt.Println("4. Менеджеры")
	fmt.Println("5. Клиенты")
	fmt.Println("------------------------")
	fmt.Println("6. Зарегистрироваться/Войти/Выйти")
	fmt.Println("------------------------")
	fmt.Println("0. Завершить работу\n")

	fmt.Print("\nВведите номер пункта: ")
	choice, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	return strings.TrimSpace(choice)
}
