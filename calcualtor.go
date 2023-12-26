package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arabicToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanNum := ""
	i := 0
	for num > 0 {
		for num >= val[i] {
			romanNum += syb[i]
			num -= val[i]
		}
		i++
	}
	return romanNum
}

func romanToArabic(s string) int {
	romanDict := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanDict[s[i:i+2]] > 0 {
			result += romanDict[s[i:i+2]]
			i++
		} else {
			result += romanDict[string(s[i])]
		}
	}
	return result
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("недопустимая арифметическая операция")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите выражение (например, 5 + 3): ")
	scanner.Scan()
	expression := scanner.Text()

	parts := strings.Fields(expression)
	if len(parts) != 3 {
		fmt.Println("Ошибка: неверный формат выражения")
		os.Exit(1)
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	aStrUpper, bStrUpper := strings.ToUpper(aStr), strings.ToUpper(bStr)

	a, err := strconv.Atoi(aStr)
	if err != nil {
		returnA := romanToArabic(aStrUpper)
		if returnA == 0 {
			fmt.Println("Ошибка: неверные типы чисел")
			os.Exit(1)
		}
		a = returnA
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		returnB := romanToArabic(bStrUpper)
		if returnB == 0 {
			fmt.Println("Ошибка: неверные типы чисел")
			os.Exit(1)
		}
		b = returnB
	}

	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		os.Exit(1)
	}

	if aStrUpper == aStr && bStrUpper == bStr {
		resultStr := arabicToRoman(result)
		fmt.Printf("Результат: %s\n", resultStr)
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
