package main

import (
	"fmt"
	"math/rand"
)

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

func main() {
	playerHP := 20
	monsterHP := 15
	inventory := []string{"зелье", "меч"}
	cmd := ""

	for playerHP > 0 && monsterHP > 0 {
		showStatus(playerHP, monsterHP, inventory)
		fmt.Println("1 - ударить, 2 - лечиться, 3 - сдаться")
		fmt.Scan(&cmd)

		switch cmd {
		case "1":
			damage := 5
			if hasItem(inventory, "меч") {
				damage = 8
				fmt.Println("Удар мечом! Урон:", damage)
			} else {
				fmt.Println("Удар кулаком! Урон:", damage)
			}
			monsterHP = hit(monsterHP, damage)

			monsterDamage := rand.Intn(4) + 1 // 1..4
			playerHP = hit(playerHP, monsterDamage)
			fmt.Println("Монстр ответил! Урон:", monsterDamage)
		case "2":
			if hasItem(inventory, "зелье") {
				inventory = removeItem(inventory, "зелье")
				playerHP = playerHP + 8
				fmt.Println("Выпил зелье! Теперь у тебя", playerHP, "HP.")
			} else {
				fmt.Println("Зелий нет.")
			}
		case "3":
			playerHP = 0
		}
	}

	showStatus(playerHP, monsterHP, inventory)
	if playerHP > 0 {
		fmt.Println("Победа!")
	} else {
		fmt.Println("Поражение.")
	}
}
