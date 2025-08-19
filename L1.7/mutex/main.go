package main

import (
	"fmt"
	"sync"
)

//Реализовать безопасную для конкуренции запись данных в структуру map.
//Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
//Проверьте работу кода на гонки (util go run -race).

// Worker пишет данные в map с защитой мьютексом
func worker(m map[int]int, wg *sync.WaitGroup, mu *sync.Mutex, id int) {
	// Гарантируем, что при завершении горутины уменьшится счетчик WaitGroup
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		// Блокируем доступ к map, чтобы избежать гонки данных
		mu.Lock()
		// Записываем данные в map по ключу id
		m[id] = i
		// Разблокируем доступ к map
		mu.Unlock()
	}
}

func main() {
	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Создаем мьютекс для синхронизации доступа к map
	var mu sync.Mutex

	// Инициализируем обычный map
	m := make(map[int]int, 3)

	// Запускаем три горутины worker
	for i := 1; i <= 3; i++ {
		// Увеличиваем счетчик горутин в WaitGroup
		wg.Add(1)
		// Запускаем worker в отдельной горутине
		go worker(m, &wg, &mu, i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим итоговое содержимое мапы
	fmt.Println(m)
}
