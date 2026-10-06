package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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
