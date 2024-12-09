package main

import "fmt"

func main() {
	numbers := []int{5, 2, 8, 1, 9}
	min := numbers[0] // Предполагаем, что первый элемент минимален

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	fmt.Println("Минимальное значение:", min) // Вывод: Минимальное значение: 1
}
