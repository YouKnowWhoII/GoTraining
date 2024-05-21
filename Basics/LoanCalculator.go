package main

import (
	"fmt"
	"math"
)

func calculateLoanPayment() {
	var principal float64
	var interestRate float64
	var loanTerm float64

	fmt.Print("Enter the principal amount: ")
	fmt.Scanln(&principal)
	fmt.Print("Enter the annual interest rate (as a decimal, e.g., 0.05 for 5%): ")
	fmt.Scanln(&interestRate)
	fmt.Print("Enter the loan term (in months): ")
	fmt.Scanln(&loanTerm)

	monthlyInterestRate := interestRate / 12.0
	monthlyPayment := principal * monthlyInterestRate / (1 - math.Pow(1+monthlyInterestRate, -loanTerm))

	fmt.Printf("The monthly loan payment is: %.2f\n", monthlyPayment)
}
