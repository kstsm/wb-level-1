package main

import "fmt"

//Разработать программу, которая в runtime способна определить тип переменной,
//переданной в неё (на вход подаётся interface{}).
//Типы, которые нужно распознавать: int, string, bool, chan (канал).
//Подсказка: оператор типа switch v.(type) поможет в решении.

// Функция для определения типа переменной в runtime
func detectType(v interface{}) {
	// Type switch позволяет проверить конкретный тип переменной, хранящейся в interface{}
	switch t := v.(type) {
	case int:
		// Если переменная имеет тип int
		fmt.Println("Тип: int, значение:", t)
	case string:
		// Если переменная имеет тип string
		fmt.Println("Тип: string, значение:", t)
	case bool:
		// Если переменная имеет тип bool
		fmt.Println("Тип: bool, значение:", t)
	case chan int:
		// Если переменная является каналом int
		fmt.Println("Тип: chan int, значение:", t)
	default:
		// Для всех остальных типов
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	// Примеры переменных разных типов
	a := 42
	b := "hello"
	c := true
	d := make(chan int)

	// Определяем типы переменных через функцию detectType
	detectType(a)
	detectType(b)
	detectType(c)
	detectType(d)
}
