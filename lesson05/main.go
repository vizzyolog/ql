package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Player struct {
	HP        int
	Inventory []string
}

type Monster struct {
	Name   string
	HP     int
	Damage [2]int // минимум и максимум урона
}

func (p *Player) hit(damage int) {
	p.HP -= damage
	if p.HP < 0 {
		p.HP = 0
	}
}

func (p *Player) hasItem(name string) bool {
	for _, item := range p.Inventory {
		if item == name {
			return true
		}
	}
	return false
}

func (p *Player) addItem(name string) {
	if !p.hasItem(name) {
		p.Inventory = append(p.Inventory, name)
	}
}

func (p *Player) removeItem(name string) {
	result := []string{}
	removed := false
	for _, item := range p.Inventory {
		if item == name && !removed {
			removed = true
			continue
		}
		result = append(result, item)
	}
	p.Inventory = result
}

func (m *Monster) hit(damage int) {
	m.HP -= damage
	if m.HP < 0 {
		m.HP = 0
	}
}

func (m Monster) attack() int {
	return rand.Intn(m.Damage[1]-m.Damage[0]+1) + m.Damage[0]
}

func ask() string {
	answer := ""
	fmt.Print("> ")
	fmt.Scan(&answer)
	return answer
}

func saveGame(p *Player) {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Не удалось преобразовать сохранение:", err)
		return
	}
	err = os.WriteFile("save.json", data, 0644)
	if err != nil {
		fmt.Println("Не удалось записать файл:", err)
		return
	}
	fmt.Println("Игра сохранена.")
}

func loadGame() *Player {
	data, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println("Файл сохранения не найден.")
		return nil
	}

	p := &Player{}
	err = json.Unmarshal(data, p)
	if err != nil || p.HP < 0 {
		fmt.Println("Файл сохранения повреждён.")
		return nil
	}

	fmt.Println("Сохранение загружено. HP:", p.HP, "Инвентарь:", p.Inventory)
	return p
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
