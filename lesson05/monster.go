package main

import "math/rand"

type Monster struct {
	Name   string
	HP     int
	Damage [2]int // минимум и максимум урона
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
