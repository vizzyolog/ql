package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Player struct {
	Name      string
	HP        int
	Inventory []string
}

// Must останавливает сервер при запуске, если шаблон не удалось прочитать.
var page = template.Must(template.ParseFiles("index.html"))

func game(w http.ResponseWriter, r *http.Request) {
	player := Player{
		Name:      "Алёша",
		HP:        15,
		Inventory: []string{"Меч", "Зелье", "Щит"},
	}

	// Execute подставляет поля player в шаблон и отправляет HTML браузеру.
	err := page.Execute(w, player)
	if err != nil {
		http.Error(w, "Не удалось показать страницу", http.StatusInternalServerError)
	}
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

// Запуск из папки lesson11:
// go run main.go
//
// Чтобы остановить сервер, нажми Ctrl+C.
