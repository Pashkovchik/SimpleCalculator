package main

import (
	Calculate "TestProject/pkg"
	"fmt"
)

func main() {
	var a, b float32
	var s string
	fmt.Println("Введите первое число")
	fmt.Scanln(&a)
	fmt.Println("Введите действие")
	fmt.Scanln(&s)
	fmt.Println("Введите второе число")
	fmt.Scanln(&b)
	fmt.Println(Calculate.Calc(a, b, s))
}
