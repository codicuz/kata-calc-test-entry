package main

import (
	"errors"
	"strconv"
)

var romanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
}

var arabicNumerals = []struct {
	Value  int
	Symbol string
}{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(s string) (int, error) {
	total := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		curr, exists := romanNumerals[string(s[i])]
		if !exists {
			return 0, errors.New("Не правильная риская цифра")
		}

		if curr < prev {
			total -= curr
		} else {
			total += curr
		}
		prev = curr
	}
	return total, nil
}

func intToRoman(num int) string {
	result := ""

	for _, numeral := range arabicNumerals {
		for num >= numeral.Value {
			result += numeral.Symbol
			num -= numeral.Value
		}
	}

	return result
}

func convertRomanToInt(input string) (int, error) {
	return romanToInt(input)
}

func convertIntToRoman(num int) string {
	return intToRoman(num)
}

func isRoman(s string) bool {
	_, err := romanToInt(s)
	return err == nil
}

func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isValidRoman(s string) (bool, error) {
	validRomans := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, roman := range validRomans {
		if s == roman {
			return true, nil
		}
	}
	return false, errors.New("Не правильное римское число")
}
