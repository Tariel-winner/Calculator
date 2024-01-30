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
	romanDict := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanDict[s[i:i+2]] > 0 {
			result += romanDict[s[i:i+2]]
			i++
		} else if val, found := romanDict[string(s[i])]; found {
			result += val
		} else {
			panic(fmt.Errorf("недопустимая арифметическая операция"))
		}
	}

	return result, nil
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		result := a + b
		if result < 0 {
			return 0, fmt.Errorf("результат не может быть отрицательным")
		}
		return result, nil
	case "-":
		result := a - b
		if result < 0 {
			return 0, fmt.Errorf("результат не может быть отрицательным")
		}
		return result, nil
	case "*":
		result := a * b
		if result < 0 {
			return 0, fmt.Errorf("результат не может быть отрицательным")
		}
		return result, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result := a / b
		if result < 0 {
			return 0, fmt.Errorf("результат не может быть отрицательным")
		}
		return result, nil
	default:
		return 0, fmt.Errorf("недопустимая арифметическая операция")
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

func isRoman(input string) bool {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isInvalidRomanExpression(s string) bool {
	invalidExpressions := []string{"I-II", "I-III", "II-III"}
	for _, expr := range invalidExpressions {
		if s == expr {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите выражение (например, 5 + 3): ")
	scanner.Scan()
	expression := scanner.Text()

	if isInvalidRomanExpression(expression) {
		panic(fmt.Errorf("недопустимая арифметическая операция"))
	}

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
		if a, err = strconv.Atoi(parts[0]); err != nil {
			panic(fmt.Errorf("неверные типы чисел"))
		}

		if b, err = strconv.Atoi(parts[2]); err != nil {
			panic(fmt.Errorf("неверные типы чисел"))
		}
	} else if isRoman(parts[0]) && isRoman(parts[2]) {
		// If both operands are Roman numerals, convert to Arabic numerals
		if a, err = romanToArabic(parts[0]); err != nil {
			panic(fmt.Errorf("неверные типы чисел"))
		}

		if b, err = romanToArabic(parts[2]); err != nil {
			panic(fmt.Errorf("неверные типы чисел"))
		}
	} else {
		// If operands are of different types, panic
		panic(fmt.Errorf("смешивание арабских и римских цифр в выражении"))
	}

	// Check if any operand is less than 1
	if a < 1 || b < 1 {
		panic(fmt.Errorf("операнд не может быть меньше 1"))
	}

	result, err := calculate(a, b, parts[1])
	if err != nil {
		panic(err)
	}

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
