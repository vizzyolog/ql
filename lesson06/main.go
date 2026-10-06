package main

import "fmt"

// Hero — карточка игрового персонажа.
// struct собирает разные данные в один «пакет».
type Hero struct {
	Name  string
	HP    int
	Level int
}

func main() {
	// Создаём героя — заполняем поля по именам
	hero := Hero{Name: "Алёша", HP: 20, Level: 1}
	fmt.Println(hero)

	// Читаем одно поле через точку
	fmt.Println("Зовут:", hero.Name)
	fmt.Println("Здоровье:", hero.HP)

	// Меняем поле — как обычную переменную
	hero.HP = hero.HP - 5
	fmt.Println("После удара:", hero.HP)

	// Второй герой — те же поля, своё значение
	// Поля, которые не указали, получают «нулевые» значения
	enemy := Hero{Name: "Гоблин"}
	fmt.Println(enemy)
}
