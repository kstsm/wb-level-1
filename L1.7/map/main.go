package main

import (
	"fmt"
	"sync"
)

//Реализовать безопасную для конкуренции запись данных в структуру map.
//Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
//Проверьте работу кода на гонки (util go run -race).

// Worker записывает данные в sync.Map
func worker(wg *sync.WaitGroup, sm *sync.Map, id int) {
	// Гарантируем, что счетчик WaitGroup уменьшится после завершения горутины
	defer wg.Done()

	// Цикл записывает значения от 1 до 5 для ключа id
	for i := 1; i <= 5; i++ {
		// запись в sync.Map ключа и значения
		sm.Store(id, i)
	}
}

func main() {
	// Создаем потокобезопасный sync.Map
	var sm sync.Map
	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем три worker-горутины
	for i := 1; i <= 3; i++ {
		// Увеличиваем счетчик WaitGroup
		wg.Add(1)
		// Запускаем worker в отдельной горутине
		go worker(&wg, &sm, i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Обходим все элементы sync.Map и выводим их
	sm.Range(func(k, v any) bool {
		fmt.Printf("key=%v value=%v\n", k, v)
		// true позволяет продолжить обход всех элементов
		return true
	})
}
