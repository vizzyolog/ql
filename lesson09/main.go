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
	// Каждый бой — это пара: кто дерётся и с кем.
	// Один срез вместо двух — и никаких индексов.
	type Battle struct {
		hero      Hero
		monster   string
		monsterHP int
	}

	battles := []Battle{
		{hero: Hero{Name: "Алёша", HP: 20, Damage: 5}, monster: "гоблин", monsterHP: 10},
		{hero: Hero{Name: "Олег", HP: 25, Damage: 3}, monster: "дракон", monsterHP: 20},
		{hero: Hero{Name: "Маша", HP: 18, Damage: 6}, monster: "тролль", monsterHP: 15},
	}

	// Буфер — от len(battles): отправка никогда не ждёт читателя.
	result := make(chan string, len(battles))

	// Три настоящие горутины: все бои идут ПАРАЛЛЕЛЬНО.
	// Перебираем бои напрямую — без всяких индексов.
	for _, battle := range battles {
		go battle.hero.fight(battle.monster, battle.monsterHP, result)
	}

	// Будильник: если за 10 секунд все не закончат — отменяем гонку.
	// time.After запускается сразу, но сообщение придёт только через 10 секунд.
	deadline := time.After(10 * time.Second)

	// select ждёт сразу два события: результат боя или звонок будильника.
	// Какой канал пришлёт первым — та ветка и выполнится.
	// Побеждает тот, чьё сообщение пришло первым — но мы дослушаем всех.
	finished := 0
	for finished < len(battles) {
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
