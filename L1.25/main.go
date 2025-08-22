package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
//которая приостанавливает выполнение текущей горутины.
//Важно: в отличии от настоящей time.Sleep, ваша функция должна именно блокировать выполнение
//(например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.
//Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).

// time.After создаёт таймер и возвращает канал, в который автоматически
// придёт сигнал по истечении указанного времени.
// Чтение <-time.After(duration) блокирует горутину до получения сигнала.
func sleepAfter(duration time.Duration) {
	fmt.Println("Горутина остановилась функцией sleepAfter")
	<-time.After(duration)
}

// Внутри создаётся таймер, который через duration отправляет сигнал в свой канал C.
// Чтение <-timer.C блокирует горутину до момента срабатывания таймера.
func sleepTimer(duration time.Duration) {
	fmt.Println("Горутина остановилась функцией sleepTimer")
	timer := time.NewTimer(duration)
	// Останавливаем таймер, чтобы избежать утечки памяти
	defer timer.Stop()
	<-timer.C
}

// Горутина активно проверяет прошло ли нужное время
// Этот подход очень сильно нагружает CPU, он годится только для учебных целей
func sleepLoop(duration time.Duration) {
	fmt.Println("Горутина остановилась функцией sleepLoop")
	start := time.Now()
	for time.Since(start) < duration {
		// Пустой цикл, ждем пока не пройдет duration
	}
}

func main() {
	var wg sync.WaitGroup

	pause := 2 * time.Second

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 4; i++ {
			fmt.Println("Горутина работает")

			// Поочередно проверяем функции
			switch {
			case i == 0:
				sleepAfter(pause)
			case i == 1:
				sleepLoop(pause)
			case i == 2:
				sleepTimer(pause)
			}
		}
	}()

	wg.Wait()
}
