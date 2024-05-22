package main

import "fmt"

func calculate() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", add(num1, num2))
	case "-":
		fmt.Printf("Result: %.2f\n", subtract(num1, num2))
	case "*":
		fmt.Printf("Result: %.2f\n", multiply(num1, num2))
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", divide(num1, num2))
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Error: Invalid operator.")
	}
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}
