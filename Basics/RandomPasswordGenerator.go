package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomPassword() {
	var length int
	fmt.Print("Enter the desired password length: ")
	fmt.Scanln(&length)

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"!@#$%^&*()")

	var password string
	for i := 0; i < length; i++ {
		password += string(chars[rand.Intn(len(chars))])
	}

	fmt.Println("Your random password is:", password)
}
