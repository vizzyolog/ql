package main

import (
	"fmt"
	"math/rand"
)

func ask() string {
	answer := ""
	fmt.Print("> ")
	fmt.Scan(&answer)
	return answer
}

func showStatus(playerHP int, monsterHP int, inv []string) {
	fmt.Println("У тебя", playerHP, "HP. У монстра", monsterHP, "HP.")
	fmt.Println("Инвентарь:", inv)
}

func hit(hp int, damage int) int {
	hp = hp - damage
	if hp < 0 {
		hp = 0
	}
	return hp
}

func hasItem(inv []string, name string) bool {
	for _, item := range inv {
		if item == name {
			return true
		}
	}
	return false
}

func removeItem(inv []string, name string) []string {
	result := []string{}
	removed := false
	for _, item := range inv {
		if item == name && !removed {
			removed = true
			continue
		}
		result = append(result, item)
	}
	return result
}

func doStart() string {
	fmt.Println("Ты в тёмном коридоре. 1 - налево (сундук), 2 - направо (дракон)")
	answer := ask()
	if answer == "1" {
		return "left"
	}
	if answer == "2" {
		return "right"
	}
	return "start"
}

func doLeft(inv []string) []string {
	fmt.Println("Комната с сундуком. 1 - открыть, 2 - вернуться")
	answer := ask()
	if answer == "1" {
		if !hasItem(inv, "зелье") {
			inv = append(inv, "зелье")
			fmt.Println("Ты нашёл зелье!")
		}
		if !hasItem(inv, "меч") {
			inv = append(inv, "меч")
			fmt.Println("Ты нашёл меч!")
		}
	}
	return inv
}

func doRight() string {
	fmt.Println("Логово дракона. 1 - убежать домой, 2 - вступить в бой")
	answer := ask()
	if answer == "1" {
		return "end"
	}
	if answer == "2" {
		return "fight"
	}
	return "right"
}

// Весь бой до победы или поражения. HP живут внутри функции.
func doFight(inv []string) []string {
	playerHP := 20
	monsterHP := 15

	for playerHP > 0 && monsterHP > 0 {
		showStatus(playerHP, monsterHP, inv)
		fmt.Println("1 - ударить, 2 - лечиться, 3 - сдаться")
		cmd := ask()

		switch cmd {
		case "1":
			damage := 5
			if hasItem(inv, "меч") {
				damage = 8
				fmt.Println("Удар мечом! Урон:", damage)
			} else {
				fmt.Println("Удар кулаком! Урон:", damage)
			}
			monsterHP = hit(monsterHP, damage)

			monsterDamage := rand.Intn(4) + 1
			playerHP = hit(playerHP, monsterDamage)
			fmt.Println("Дракон ответил! Урон:", monsterDamage)
		case "2":
			if hasItem(inv, "зелье") {
				inv = removeItem(inv, "зелье")
				playerHP = playerHP + 8
				fmt.Println("Выпил зелье! Теперь у тебя", playerHP, "HP.")
			} else {
				fmt.Println("Зелий нет.")
			}
		case "3":
			playerHP = 0
		}
	}

	if playerHP > 0 {
		fmt.Println("Победа над драконом!")
	} else {
		fmt.Println("Поражение.")
	}
	return inv
}

func main() {
	room := "start"
	inventory := []string{}

	for room != "end" {
		switch room {
		case "start":
			room = doStart()
		case "left":
			inventory = doLeft(inventory)
			room = "start"
		case "right":
			room = doRight()
		case "fight":
			inventory = doFight(inventory)
			room = "end"
		}
	}

	fmt.Println("Игра окончена. Инвентарь:", inventory)
}
