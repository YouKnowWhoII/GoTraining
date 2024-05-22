package main

import (
	"fmt"
	"os"
)

func greetUser() {
	fmt.Println("Hello, World!")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	greetUser()
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		calculateArea()
	case 2:
		calculate()
	case 3:
		playGuessingGame()
	case 4:
		convertTemperature()
	case 5:
		calculateLoanPayment()
	case 6:
		analyzeText()
	case 7:
		startWebServer()
	case 8:
		checkPasswordStrength()
	case 9:
		generateRandomPassword()
	case 10:
		processFile()
	default:
		fmt.Println("Invalid choice")
		os.Exit(1)
	}
}
