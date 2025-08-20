package main

import "fmt"

//Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
//Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
//Для выбора опорного элемента можно взять середину или первый элемент.

func quickSort(arr []int) []int {
	// Если в массиве меньше 2 элементов сортировка прекращается
	// без неё сортировка никогда бы не остановилась
	if len(arr) < 2 {
		return arr
	}

	// Берём опорный элемент (середину массива),
	// чтобы уменьшить вероятно потери скорости в сортировке
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Два массива:
	// less - числа меньше или равные опорному
	// greater - числа больше опорного
	var less, greater []int

	// Проходим по всем элементам, кроме самого pivot
	for i, num := range arr {
		// Пропускаем сам pivot, чтобы не сравнивать его с самим собой
		if i == pivotIndex {
			continue
		}
		// Если элемент меньше либо равен pivot идет less
		if num <= pivot {
			less = append(less, num)
		} else {
			// Если элемент больше pivot идет в greater
			greater = append(greater, num)
		}
	}

	// Сначала сортируем всё, что меньше pivot,
	// потом добавляем сам pivot,
	// а потом сортируем всё, что больше pivot
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)

	// Возвращаем собранный отсортированный массив
	return result
}

func main() {
	arr := []int{2, 1, 6, 8, 5, 7, 2}

	res := quickSort(arr)
	fmt.Println(res) // Вывод [1 2 2 5 6 7 8]

}
