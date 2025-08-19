package main

import (
	"fmt"
	"sync"
	"time"
)

//Остановка горутины через выход по условию

func worker(t int, wg *sync.WaitGroup) {
	// Гарантируем, что Done будет вызван в конце работы горутины
	defer wg.Done()

	// Цикл выполняется t раз
	for i := 0; i < t; i++ {
		// Сообщение о работе горутины
		fmt.Println("Горутина работает")
		// Пауза между итерациями
		time.Sleep(time.Second)
	}

	// Сообщение о завершении горутины
	fmt.Println("Горутина завершена по условию")
}

func main() {
	// Создаем WaitGroup для ожидания завершения горутины
	var wg sync.WaitGroup

	// Увеличиваем счетчик на 1 перед запуском горутины
	wg.Add(1)
	// Запускаем горутину
	go worker(5, &wg)

	// Ждем завершения горутины в WaitGroup
	wg.Wait()
}
