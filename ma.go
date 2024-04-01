package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	MIN = 1
	MAX = 10
)

var romanOutput = map[int]string{
	1: "I", 4: "IV", 5: "V", 9: "IX",
	10: "X", 40: "XL", 50: "L", 90: "XC",
	100: "C",
}

func toRoman(r int) string {
	var keys []int
	for k := range romanOutput {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var result string
	num := r
	for _, value := range keys {
		for num >= value {
			result += romanOutput[value]
			num -= value
		}
	}

	return result
}

var roman = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func isRomanNumber(s string) bool {
	_, exists := roman[s]
	return exists
}

func toArabic(n string) int {
	if val, ok := roman[n]; ok {
		return val
	}
	panic("Неверный ввод: введено не римское число")
}

func parseExpression(expr string) (string, string, string) {
	expr = strings.TrimSpace(strings.ReplaceAll(expr, " ", ""))
	for _, op := range []string{"+", "-", "*", "/"} {
		if strings.Contains(expr, op) {
			parts := strings.Split(expr, op)
			if len(parts) == 2 {
				return parts[0], parts[1], op
			} else {
				panic("некорректное выражение")
			}
		}
	}
	panic("операция не найдена")
}

func validateAndConvert(a, b string) (int, int, bool) {
	isRoman := isRomanNumber(a) && isRomanNumber(b)
	var aAsInt, bAsInt int

	if isRoman {
		aAsInt = toArabic(a)
		bAsInt = toArabic(b)
	} else {
		aAsIntC, errA := strconv.Atoi(a)
		bAsIntC, errB := strconv.Atoi(b)
		if errA == nil && errB == nil {
			aAsInt = aAsIntC
			bAsInt = bAsIntC
		}
	}
	if MIN > aAsInt || aAsInt > MAX || MIN > bAsInt || bAsInt > MAX {
		panic("Числа должны быть от 1 до 10")
	}
	return aAsInt, bAsInt, isRoman
}

func calculate(a, b int, op string, isRoman bool) int {
	switch op {
	case "+":
		return a + b
	case "-":
		result := a - b
		if result < 1 && isRoman {
			panic("Результат вычитания римских цифр меньше 1")
		}
		return result
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}
	}
	panic("Подходящая операция не найдена")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение")
		text, _ := reader.ReadString('\n')
		number1, number2, op := parseExpression(text)
		a, b, isRoman := validateAndConvert(number1, number2)
		result := calculate(a, b, op, isRoman)
		if isRoman {
			fmt.Println("Результат: ", toRoman(result))
		} else {
			fmt.Println("Результат: ", result)
		}
	}
}
