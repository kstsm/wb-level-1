package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Human - базовая структура с произвольными полями и методами
type Human struct {
	Name string
	Age  int
}

// Action - структура, которая встраивает Human
// Благодаря встраиванию Action получает доступ к полям и методам Human
type Action struct {
	Human // встраиваем Human
	Sport string
}

// SayHello выводит информацию о человеке
func (h Human) SayHello() string {
	return fmt.Sprintf("Привет! Меня зовут %s, мне %d лет", h.Name, h.Age)
}

// SayHobby выводит информацию об увлечении.
// Благодаря встраиванию можно использовать поля из Human.
func (a Action) SayHobby() string {
	return fmt.Sprintf("Меня зовут %s, мне %d лет и я люблю %s", a.Name, a.Age, a.Sport)
}

func main() {
	// Создаём объект bob из структуры Action
	// Заполняем данные как для Human и Action
	bob := Action{
		Human: Human{
			Name: "Боб",
			Age:  23,
		},
		Sport: "теннис",
	}

	// Вызываем оба метода и печатаем их
	fmt.Println(bob.SayHello())
	fmt.Println(bob.SayHobby())
}
