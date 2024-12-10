package main

import "fmt"

// Функция для быстрой сортировки массива
func quicksort(arr []int) []int {
	// Если массив состоит из одного элемента или пустой, возвращаем его
	if len(arr) <= 1 {
		return arr
	}

	// Выбираем опорный элемент (например, первый элемент)
	pivot := arr[0]
	left := []int{}
	right := []int{}
	middle := []int{pivot}

	// Разбиваем массив на элементы меньше, равные и больше опорного элемента
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}

	// Рекурсивно сортируем левую и правую части массива
	left = quicksort(left)
	right = quicksort(right)

	// Объединяем отсортированные части
	return append(append(left, middle...), right...)
}

func main() {
	// Пример использования
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Original array:", arr)

	sortedArr := quicksort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
