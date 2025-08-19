package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Остановка горутины через runtime.Goexit

// Worker выводит числа и завершает себя с помощью runtime.Goexit()
func worker(wg *sync.WaitGroup) {
	// Сообщаем WaitGroup о завершении горутины после выхода из функции
	//(оно не сработает т.к. runtime.Goexit() быстрей закроет горутину)
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// Вывод текущего числа
		fmt.Println(i)

		// Если достигнут последний индекс, выводим сообщение и завершаем горутину
		if i == 3 {
			fmt.Println("Завершение горутины", i)
			// Немедленно завершаем текущую горутину
			runtime.Goexit()
		}

		// Пауза между итерациями
		time.Sleep(time.Second)
	}
}

func main() {
	// Создаем WaitGroup для ожидания завершения горутины
	var wg sync.WaitGroup

	// Увеличиваем счетчик на 1, так как запускаем одну горутину
	wg.Add(1)

	// Запускаем воркер в отдельной горутине
	go worker(&wg)

	// Ждем завершения всех горутин в WaitGroup
	wg.Wait()
}
