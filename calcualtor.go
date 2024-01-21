package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func romanToArabic(s string) (int, error) {
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
	return result, nil
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		result := a - b
		if result <= 0 {
			panic(fmt.Errorf("результат меньше или равен нулю"))
		}
		return result
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic(fmt.Errorf("деление на ноль"))
		}
		result := a / b
		if result <= 0 {
			panic(fmt.Errorf("результат меньше или равен нулю"))
		}
		return result
	default:
		panic(fmt.Errorf("недопустимая арифметическая операция"))
	}
}

func isArabic(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите выражение (например, 5 + 3): ")
	scanner.Scan()
	expression := scanner.Text()

	var parts []string

	// Check if the expression contains spaces
	if strings.Contains(expression, " ") {
		// Split the expression into parts using spaces
		parts = strings.Fields(expression)
	} else {
		// Manually extract operands and operator
		parts = []string{string(expression[0]), string(expression[1]), string(expression[2])}
	}

	if len(parts) != 3 {
		panic(fmt.Errorf("неверный формат выражения"))
	}

	var a, b int
	var err error

	// Check if both operands are Arabic numerals
	if isArabic(parts[0]) && isArabic(parts[2]) {
		if a, err = strconv.Atoi(parts[0]); err != nil || a < 1 || a > 10 {
			panic(fmt.Errorf("неверные типы чисел"))
		}

		if b, err = strconv.Atoi(parts[2]); err != nil || b < 1 || b > 10 {
			panic(fmt.Errorf("неверные типы чисел"))
		}
	} else {
		// If not Arabic numerals, try to convert to Roman numerals
		if a, err = romanToArabic(parts[0]); err != nil {
			panic(fmt.Errorf("неверные типы чисел"))
		}

		if b, err = romanToArabic(parts[2]); err != nil {
			panic(fmt.Errorf("неверные типы чисел"))
		}
	}

	result := calculate(a, b, parts[1])

	// Check if the result should be displayed as Arabic or Roman numeral
	if isArabic(parts[0]) && isArabic(parts[2]) {
		// If both operands are Arabic numerals, display the result as Arabic numeral
		fmt.Printf("Результат: %d\n", result)
	} else {
		// If operands are Roman numerals, display the result as Roman numeral
		resultStr := arabicToRoman(result)
		fmt.Printf("Результат: %s\n", resultStr)
	}
}
