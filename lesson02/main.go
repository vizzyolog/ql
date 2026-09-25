package main

import "fmt"

func main() {
	room := "start"
	hasKey := false
	answer := ""
	for room != "end" {
		switch room {
		case "start":
			fmt.Println("Ты в тёмном коридоре. 1 - налево, 2 - направо")
			fmt.Scan(&answer)
			if answer == "1" {
				room = "left"
			}
			if answer == "2" {
				room = "right"
			}
		case "left":
			fmt.Println("Комната с сундуком. 1 - открыть, 2 - вернуться")
			fmt.Scan(&answer)
			if answer == "1" {
				fmt.Println("Ты нашёл ключ!")
				hasKey = true
			}
			room = "start"
		case "right":
			fmt.Println("Здесь дракон! 1 - бежать, 2 - драться")
			fmt.Scan(&answer)
			if answer == "2" && hasKey {
				fmt.Println("Ключ оказался мечом. Победа!")
			} else {
				fmt.Println("Дракон тебя съел.")
			}
			room = "end"
		}
	}
	fmt.Println("Игра окончена.")
}
