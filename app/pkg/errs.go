package pkg

import "fmt"

func ProcessErrors(err error) {
	if err != nil {
		fmt.Printf("\nОшибка: %v\n", err.Error())
	}
}
