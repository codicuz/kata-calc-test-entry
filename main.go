package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: 3 + 5 или IV * II):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		panic("Не верный оператор")
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		panic("Не верный формат математической операции")
	}

	aStr, bStr := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	var a, b int
	var err error
	var result int

	if isRoman(aStr) && isRoman(bStr) {
		_, err = isValidRoman(aStr)
		if err != nil {
			panic(err)
		}

		_, err = isValidRoman(bStr)
		if err != nil {
			panic(err)
		}

		a, err = convertRomanToInt(aStr)
		if err != nil {
			panic(err)
		}

		b, err = convertRomanToInt(bStr)
		if err != nil {
			panic(err)
		}
	} else if isArabic(aStr) && isArabic(bStr) {
		a, err = strconv.Atoi(aStr)
		if err != nil {
			panic(err)
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Нельзя смешивать римские и арабские числа")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10")
	}

	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		panic("Не верный оператор")
	}

	if isRoman(aStr) {
		if result < 1 {
			panic("Результат в римском исполнении ниже чем I")
		}
		fmt.Printf("Результат: %s\n", convertIntToRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
