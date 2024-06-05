package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/payment", paymentHandler) // Установка обработчика для пути "/payment"

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запуск сервера
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(1000) != 333 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Оплата прошла успешно")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Оплата не прошла")
	}
}
