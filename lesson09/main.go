package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Hero и fight — уже знакомы из lesson08: бой идёт в фоне
// и отправляет результат в канал, когда закончится.
type Hero struct {
	Name   string
	HP     int
	Damage int
}

func (h *Hero) fight(monsterName string, monsterHP int, result chan string) {
	for monsterHP > 0 && h.HP > 0 {
		monsterHP = monsterHP - h.Damage
		h.HP = h.HP - rand.Intn(4)
		time.Sleep(300 * time.Millisecond)
	}
	if h.HP > 0 {
		result <- h.Name + " победил " + monsterName + "!"
	} else {
		result <- h.Name + " пал в бою с " + monsterName + "..."
	}
}

func main() {
	heroes := []Hero{
		{Name: "Алёша", HP: 20, Damage: 5},
		{Name: "Олег", HP: 25, Damage: 3},
		{Name: "Маша", HP: 18, Damage: 6},
	}

	monsters := []struct {
		name string
		hp   int
	}{
		{"гоблин", 10},
		{"дракон", 20},
		{"тролль", 15},
	}

	// Буфер — от len(heroes), как в lesson08: отправка никогда не ждёт читателя.
	result := make(chan string, len(heroes))

	// Три настоящие горутины: все бои идут ПАРАЛЛЕЛЬНО.
	// heroes[i] дерётся с monsters[i]: первый с первым, второй со вторым.
	for i := range heroes {
		go heroes[i].fight(monsters[i].name, monsters[i].hp, result)
	}

	// Будильник: если за 10 секунд все не закончат — отменяем гонку.
	// time.After запускается сразу, но сообщение придёт только через 10 секунд.
	deadline := time.After(10 * time.Second)

	// select ждёт сразу два события: результат боя или звонок будильника.
	// Какой канал пришлёт первым — та ветка и выполнится.
	// Побеждает тот, чьё сообщение пришло первым — но мы дослушаем всех.
	finished := 0
	for finished < len(heroes) {
		select {
		case msg := <-result:
			finished++
			if finished == 1 {
				fmt.Println("ПЕРВЫМ закончил:", msg)
			} else {
				fmt.Println(msg)
			}
		case <-deadline:
			fmt.Println("Время вышло! Кто ещё в бою — отступает.")
			return
		}
	}

	fmt.Println("Все бои окончены.")

	// --- Маленькая практика ---
	// 1. Уменьшите время дедлайна (например, до 3 секунд) —
	//    некоторые бои не успеют закончиться. Что выведет программа?
	// 2. Сделайте так, чтобы после дедлайна программа выводила,
	//    сколько героев всё-таки успели (счётчик finished).
	// 3. Усложнение: пусть герой отправляет в канал своё оставшееся HP,
	//    а main награждает «лучшего бойца» — с наибольшим HP.
}
