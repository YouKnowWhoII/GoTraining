package main

import (
	"fmt"
)

func convertTemperature() {
	var temp float64
	var unit string

	fmt.Print("Enter temperature: ")
	fmt.Scanln(&temp)
	fmt.Print("Is this temperature in Celsius or Fahrenheit (C/F)? ")
	fmt.Scanln(&unit)

	switch unit {
	case "C", "c":
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temp, temp*9/5+32)
	case "F", "f":
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temp, (temp-32)*5/9)
	default:
		fmt.Println("Invalid unit. Please enter C for Celsius or F for Fahrenheit.")
	}
}
