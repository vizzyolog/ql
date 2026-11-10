package main

import (
	"fmt"
	"log"
	"net/http"
)

func game(w http.ResponseWriter, r *http.Request) {
	// Вместо обычного текста сервер отправляет браузеру HTML-файл.
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", game)

	fmt.Println("Открой в браузере: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Запуск из папки lesson10:
// go run main.go
//
// Чтобы остановить сервер, нажми Ctrl+C.
