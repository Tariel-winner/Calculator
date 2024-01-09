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
