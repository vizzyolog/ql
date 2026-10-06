package main

import "fmt"

func main() {
	room := "start"
	player := &Player{HP: 20}

	fmt.Println("1 - новая игра, 2 - продолжить")
	if ask() == "2" {
		if saved := loadGame(); saved != nil {
			player = saved
		}
	}

	for room != "end" {
		switch room {
		case "start":
			room = doStart()
		case "save":
			saveGame(player)
			room = "start"
		case "left":
			doLeft(player)
			room = "start"
		case "right":
			room = doRight()
		case "fight":
			doFight(player)
			if player.HP > 0 {
				room = "start"
			} else {
				room = "end"
			}
		}
	}

	fmt.Println("Игра окончена. HP:", player.HP, "Инвентарь:", player.Inventory)
}
