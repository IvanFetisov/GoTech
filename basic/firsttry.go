package basic

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
	default:
		fmt.Println("Ошибка: неверный оператор")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
