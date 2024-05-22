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

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"!@#$%^&*()")

	var password string
	for i := 0; i < length; i++ {
		password += string(chars[r.Intn(len(chars))])
	}

	fmt.Println("Your random password is:", password)
}
