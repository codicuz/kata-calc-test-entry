package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b != 0 {
		return a / b
	} else {
		fmt.Println("Ошибка: Деление на ноль!")
		return 0
	}
}
