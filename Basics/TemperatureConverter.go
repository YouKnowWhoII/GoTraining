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
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temp, celciusToFahrenheit(temp))
	case "F", "f":
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temp, fahrenheitToCelcius(temp))
	default:
		fmt.Println("Invalid unit. Please enter C for Celsius or F for Fahrenheit.")
	}
}

func celciusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelcius(f float64) float64 {
	return (f - 32) * 5 / 9
}
