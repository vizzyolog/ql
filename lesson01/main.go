package main

import (
	"fmt"
	"io"
	"net/http"
)

const name = "Олег"
const ip = "10.116.72.6"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Привет это", name)
}

func main() {
	http.HandleFunc("/", hello)
	go http.ListenAndServe(":8080", nil)

	response, err := http.Get("http://" + ip + ":8080")
	if err != nil {
		fmt.Println("Ошибка", err)
	}

	defer response.Body.Close()
	text, _ := io.ReadAll(response.Body)

	fmt.Println("Resp: ", string(text))
}
