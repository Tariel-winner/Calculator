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
		result := a - b
		if result <= 0 {
			return 0, fmt.Errorf("результат меньше или равен нулю")
		}
		return result, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result := a / b
		if result <= 0 {
			return 0, fmt.Errorf("результат меньше или равен нулю")
		}
		return result, nil
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

	var a, b int
	var err error

	// Проверка, что числа в диапазоне от 1 до 10 включительно
	if a, err = strconv.Atoi(parts[0]); err != nil || a < 1 || a > 10 {
		fmt.Println("Ошибка: неверные типы чисел")
		os.Exit(1)
	}

	if b, err = strconv.Atoi(parts[2]); err != nil || b < 1 || b > 10 {
		fmt.Println("Ошибка: неверные типы чисел")
		os.Exit(1)
	}

	// Проверка, что числа одного формата
	if strings.ToUpper(parts[0]) == parts[0] && strings.ToUpper(parts[2]) == parts[2] {
		// Если оба числа римские, но операция арабская, выдайте ошибку
		if _, err := strconv.Atoi(parts[1]); err == nil {
			fmt.Println("Ошибка: числа разного формата")
			os.Exit(1)
		}
	} else {
		// Если числа арабские, но операция римская, выдайте ошибку
		if _, err := strconv.Atoi(parts[1]); err != nil {
			fmt.Println("Ошибка: числа разного формата")
			os.Exit(1)
		}
	}

	result, err := calculate(a, b, parts[1])
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		os.Exit(1)
	}

	// Проверка, что результат в правильном диапазоне
	if result < 1 || result > 10 {
		fmt.Println("Ошибка: результат должен быть от 1 до 10 включительно")
		os.Exit(1)
	}

	// Проверка, какой формат ввода и вывод результата соответствующим образом
	if strings.ToUpper(parts[0]) == parts[0] && strings.ToUpper(parts[2]) == parts[2] {
		resultStr := arabicToRoman(result)
		fmt.Printf("Результат: %s\n", resultStr)
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
