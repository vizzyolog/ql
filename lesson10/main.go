package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func game(w http.ResponseWriter, r *http.Request) {
	// Вместо обычного текста сервер отправляет браузеру HTML-файл.
	http.ServeFile(w, r, "index.html")
}

func start(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Нужен POST-запрос", http.StatusMethodNotAllowed)
		return
	}

	// Браузер отправил имя героя в теле запроса — просто текстом, без JSON.
	name, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Не удалось прочитать имя", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s, приключение началось! Сервер получил POST-запрос.", string(name))
}

func main() {
	http.HandleFunc("/", game)
	http.HandleFunc("/start", start)

	fmt.Println("Открой в браузере: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Запуск из папки lesson10:
// go run main.go
//
// Чтобы остановить сервер, нажми Ctrl+C.
