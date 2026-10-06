package main

import "fmt"

func ask() string {
	answer := ""
	fmt.Print("> ")
	fmt.Scan(&answer)
	return answer
}

func doStart() string {
	fmt.Println("Ты в тёмном коридоре. 1 - налево (сундук), 2 - направо (дракон), 3 - сохранить, 4 - выйти")
	answer := ask()
	if answer == "1" {
		return "left"
	}
	if answer == "2" {
		return "right"
	}
	if answer == "3" {
		return "save"
	}
	if answer == "4" {
		return "end"
	}
	return "start"
}

func doLeft(p *Player) {
	fmt.Println("Комната с сундуком. 1 - открыть, 2 - вернуться")
	answer := ask()
	if answer == "1" {
		if !p.hasItem("зелье") {
			p.addItem("зелье")
			fmt.Println("Ты нашёл зелье!")
		}
		if !p.hasItem("меч") {
			p.addItem("меч")
			fmt.Println("Ты нашёл меч!")
		}
	}
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

func doFight(p *Player) {
	dragon := Monster{Name: "Дракон", HP: 15, Damage: [2]int{1, 4}}

	for p.HP > 0 && dragon.HP > 0 {
		fmt.Println("У тебя", p.HP, "HP. У дракона", dragon.HP, "HP.")
		fmt.Println("Инвентарь:", p.Inventory)
		fmt.Println("1 - ударить, 2 - лечиться, 3 - сдаться")
		cmd := ask()

		switch cmd {
		case "1":
			damage := 5
			if p.hasItem("меч") {
				damage = 8
				fmt.Println("Удар мечом! Урон:", damage)
			} else {
				fmt.Println("Удар кулаком! Урон:", damage)
			}
			dragon.hit(damage)

			monsterDamage := dragon.attack()
			p.hit(monsterDamage)
			fmt.Println("Дракон ответил! Урон:", monsterDamage)
		case "2":
			if p.hasItem("зелье") {
				p.removeItem("зелье")
				p.HP += 8
				fmt.Println("Выпил зелье! Теперь у тебя", p.HP, "HP.")
			} else {
				fmt.Println("Зелий нет.")
			}
		case "3":
			p.HP = 0
		}
	}

	if p.HP > 0 {
		fmt.Println("Победа над драконом!")
	} else {
		fmt.Println("Поражение.")
	}
}
