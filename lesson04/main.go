package main

import "fmt"

func ask() string {
	answer := ""
	fmt.Print("> ")
	fmt.Scan(&answer)
	return answer
}

func doStart() string {
	fmt.Println("Ты в тёмном коридоре. 1 - налево, 2 - направо")
	answer := ask()
	if answer == "1" {
		return "left"
	}
	if answer == "2" {
		return "right"
	}
	return "start"
}

func doLeft(hasKey bool) (string, bool) {
	fmt.Println("Комната с сундуком. 1 - открыть, 2 - вернуться")
	answer := ask()
	if answer == "1" {
		fmt.Println("Ты нашёл ключ!")
		hasKey = true
	}
	return "start", hasKey
}

func doRight(hasKey bool) string {
	fmt.Println("Здесь дракон! 1 - бежать, 2 - драться")
	answer := ask()
	if answer == "2" && hasKey {
		fmt.Println("Ключ оказался мечом. Победа!")
	} else {
		fmt.Println("Дракон тебя съел.")
	}
	return "end"
}

func main() {
	room := "start"
	hasKey := false

	for room != "end" {
		switch room {
		case "start":
			room = doStart()
		case "left":
			room, hasKey = doLeft(hasKey)
		case "right":
			room = doRight(hasKey)
		}
	}

	fmt.Println("Игра окончена.")
}
