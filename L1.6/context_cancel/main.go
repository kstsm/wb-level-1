package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Остановка горутины через context.WithCancel

// Worker выполняет работу до сигнала отмены контекста
func worker(ctx context.Context, wg *sync.WaitGroup) {
	// Сообщаем WaitGroup о завершении горутины
	defer wg.Done()

	for {
		// Проверяем контекст на сигнал отмены
		select {
		// Если контекст отменён
		case <-ctx.Done():
			fmt.Println("Горутина закрывается через cancel")
			// Завершаем горутину
			return
		// Если отмены нет, продолжаем работу
		default:
			fmt.Println("Горутина работает")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Создаем WaitGroup для ожидания завершения горутины
	var wg sync.WaitGroup

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	// Отложенное освобождение ресурсов контекста при выходе из main
	defer cancel()

	// Увеличиваем счетчик WaitGroup на 1
	wg.Add(1)

	// Запускаем воркер в отдельной горутине
	go worker(ctx, &wg)

	// Даем горутине поработать 3 секунды
	time.Sleep(3 * time.Second)

	// Отправляем сигнал отмены контекста для горутины
	cancel()

	// Ждем завершения горутины в WaitGroup
	wg.Wait()
}
