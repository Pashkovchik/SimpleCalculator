package Calculate

import "fmt"

func Calc(x, y float32, s string) float32 {
	fmt.Println("Провожу сложные вычисления")
	fmt.Println("Еще немного...")
	switch s {
	case "+":
		fmt.Println("Сумма:")
		return x + y
	case "-":
		fmt.Println("Разность:")
		return x - y
	case "*":
		fmt.Println("Произведение:")
		return x * y
	case "/":
		if y == 0 {
			fmt.Println("на ноль делить нельзя!")
			return 0
		}
		fmt.Println("Деление:")
		return x / y
	default:
		fmt.Println("Некорректное значение")
		return 0
	}
	return x + y
}
