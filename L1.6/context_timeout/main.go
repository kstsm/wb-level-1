package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Остановка горутины через context.WithTimeout

func worker(ctx context.Context, wg *sync.WaitGroup) {
	// Сообщаем WaitGroup, что горутина завершилась после выхода из функции
	defer wg.Done()

	for {
		// Проверка канала на сигнал завершения
		select {
		// Если контекст завершён (таймаут или отмена)
		case <-ctx.Done():
			// Выводим сообщение о завершении
			fmt.Println("Горутина закрывается")
			// Завершаем работу горутины
			return
		// Если сигнала завершения нет
		default:
			// Сообщение о работе горутины
			fmt.Println("Горутина работает")
			// Пауза между итерациями
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Создаем WaitGroup для ожидания завершения горутины
	var wg sync.WaitGroup

	// Создаем контекст с таймаутом 5 секунд
	ctx, stop := context.WithTimeout(context.Background(), 5*time.Second)
	// Отмена контекста при завершении main
	defer stop()

	// Увеличиваем счетчик WaitGroup на 1 для запуска горутины
	wg.Add(1)
	// Запускаем воркер в отдельной горутине
	go worker(ctx, &wg)

	// Ждем завершения всех горутин в WaitGroup
	wg.Wait()
}
