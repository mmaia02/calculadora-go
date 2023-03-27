package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%.2f / %.2f = %.2f", num1, num2, num1/num2)
	default:
		fmt.Println("Operador inválido!")
	}
}
