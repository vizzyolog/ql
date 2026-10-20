package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Hero уже знаком из прошлых уроков
type Hero struct {
	Name   string
	HP     int
	Damage int
}

// Метод боя. Раньше бой был мгновенным — теперь каждый удар занимает время.
// time.Sleep останавливает функцию на указанное время.
// Здесь Sleep НЕ тормозит всю программу, потому что бой запущен через go.
func (h *Hero) fight(monsterName string, monsterHP int, result chan string) {
	for monsterHP > 0 && h.HP > 0 {
		monsterHP = monsterHP - h.Damage
		h.HP = h.HP - rand.Intn(4) // монстр бьёт на 0-3 урона
		time.Sleep(300 * time.Millisecond) // один «ход» боя
		fmt.Println(h.Name, "бьёт", monsterName, "... монстру осталось", monsterHP, "HP")
	}

	if h.HP > 0 {
		result <- h.Name + " победил " + monsterName + "!"
	} else {
		result <- h.Name + " пал в бою с " + monsterName + "..."
	}
}

func main() {
	// Канал — труба, по которой горутины передают друг другу сообщения.
	// chan string — в трубе гуляют строки.
	// Без канала main не узнал бы, чем закончились бои.
	result := make(chan string)

	heroes := []Hero{
		{Name: "Алёша", HP: 20, Damage: 5},
		{Name: "Олег", HP: 25, Damage: 3},
	}

	monsters := []struct {
		name string
		hp   int
	}{
		{"гоблин", 10},
		{"дракон", 20},
	}

	// go — запустить функцию ПАРАЛЛЕЛЬНО, не дожидаясь её конца.
	// Обе битвы идут одновременно: два героя дерутся в одно время,
	// а не по очереди, как было бы без go.
	for i := range heroes {
		go heroes[i].fight(monsters[i].name, monsters[i].hp, result)
	}

	// main ждёт два сообщения из канала — по одному от каждого боя.
	// Кто закончит бой первым, того сообщение придёт первым!
	for i := 0; i < 2; i++ {
		fmt.Println(<-result)
	}

	fmt.Println("Все бои окончены.")

	// --- Маленькая практика ---
	// 1. Запустите оба боя БЕЗ go — увидите разницу (бои пойдут по очереди).
	// 2. Добавьте третьего героя и третьего монстра (не забудьте поменять цикл ожидания!).
	// 3. Усложнение: пусть герой после победы отправляет в канал
	//    оставшееся HP, а main объявляет «лучшего бойца».
}
