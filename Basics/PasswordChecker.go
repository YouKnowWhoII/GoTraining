package main

import (
	"fmt"
	"unicode"
)

func checkPasswordStrength() {
	var password string
	fmt.Print("Enter a password: ")
	fmt.Scanln(&password)

	var hasUppercase, hasLowercase, hasDigit bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if len(password) >= 8 && hasUppercase && hasLowercase && hasDigit {
		fmt.Println("Your password is strong.")
	} else {
		fmt.Println("Your password is not strong enough.")
	}
}
